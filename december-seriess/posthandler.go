package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only post method allowd", http.StatusMethodNotAllowed)
		return
	}
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Not valid json", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "The users ID : %d Name : %s and Age :%d", user.Id, user.Name, user.Age)
}
func getUser(w http.ResponseWriter, r *http.Request) {

}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/createhanlder", createHandler).Methods("POST")
	r.HandleFunc("/getUser/{id}", getUser).Methods("GET")
	http.ListenAndServe(":8080", r)
}
