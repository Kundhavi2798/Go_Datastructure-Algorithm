package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Book struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	Author        string `json:"author"`
	PublishedYear int    `json:"published_year"`
}

var books = []Book{
	{ID: 1, Title: "The Go Programming Language", Author: "Alan Donovan", PublishedYear: 2015},
	{ID: 2, Title: "Clean Code", Author: "Robert C. Martin", PublishedYear: 2008},
}

var nextID = 3

func getBooks(w http.ResponseWriter, r *http.Request) {
	author := r.URL.Query().Get("author")
	w.Header().Set("Content-Type", "application/json")

	if author != "" {
		var filtered []Book
		for _, b := range books {
			if b.Author == author {
				filtered = append(filtered, b)
			}
		}
		json.NewEncoder(w).Encode(filtered)
		return
	}

	json.NewEncoder(w).Encode(books)
}

func getBookByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for _, b := range books {
		if b.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(b)
			return
		}
	}

	http.Error(w, "Book not found", http.StatusNotFound)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	var newBook Book
	err := json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newBook.ID = nextID
	nextID++

	books = append(books, newBook)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newBook)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for i, b := range books {
		if b.ID == id {
			var updatedBook Book
			if err := json.NewDecoder(r.Body).Decode(&updatedBook); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			updatedBook.ID = id
			books[i] = updatedBook
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedBook)
			return
		}
	}
	http.Error(w, "Book not found", http.StatusNotFound)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for i, b := range books {
		if b.ID == id {
			books = append(books[:i], books[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Book not found", http.StatusNotFound)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books/{id}", getBookByID).Methods("GET")
	r.HandleFunc("/books", createBook).Methods("POST")
	r.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	fmt.Println("Server running on http://localhost:8081")
	http.ListenAndServe(":8081", r)
}
