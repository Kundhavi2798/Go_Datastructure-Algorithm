package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Message struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
}

var MessageContent []Message
var NextID int

func GetMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(MessageContent)
}
func CreateMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var msg Message
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
	}
	msg.Id = NextID
	NextID++
	MessageContent = append(MessageContent, msg)
	fmt.Println(MessageContent)
	json.NewEncoder(w).Encode(&msg)
}
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/create/Messanger", CreateMessage).Methods("POST")
	router.HandleFunc("/getMessage", GetMessage).Methods("GET")
	fmt.Println("The running server:8081")
	log.Fatal(http.ListenAndServe("localhost:8081", router))
}
