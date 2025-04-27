package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Messagescom struct {
	Id      string `json:"id"`
	Message string `json:"message"`
}

var msgdata []Messagescom

func createCommunication(w http.ResponseWriter, r *http.Request) {
	var msg Messagescom
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		http.Error(w, "Unmarshal error", http.StatusBadRequest)
		return // ✅ return to stop execution
	}
	msgdata = append(msgdata, msg)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(msg)
}
func getData(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	id := param["id"]
	fmt.Println("Requested id:", id)
	fmt.Println("Available messages:", msgdata)

	for _, item := range msgdata {
		fmt.Println("Comparing:", item.Id, id)
		if item.Id == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.Error(w, "Message not found", http.StatusNotFound) // ✅ proper status
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/create/Messages", createCommunication).Methods("POST")
	router.HandleFunc("/get/data/{id}", getData).Methods("GET")
	fmt.Println("Running local host : 8081")
	log.Fatal(http.ListenAndServe(":8081", router))

}
