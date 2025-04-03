package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type UserCreation struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var result UserCreation
	err := json.NewDecoder(r.Body).Decode(&result)
	if err != nil {
		http.Error(w, "Cant decode the error", http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"Message": "User created successfully",
		"name":    result.Name,
		"email":   result.Email,
	})
}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/createUser", CreateUser).Methods("POST")

	log.Println("The running server 8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}
