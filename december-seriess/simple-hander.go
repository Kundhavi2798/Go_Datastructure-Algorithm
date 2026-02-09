package main

import (
	"fmt"
	"net/http"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}
func main() {
	http.HandleFunc("/get", Welcome)
	fmt.Println("The server running 9090")
	http.ListenAndServe(":9090", nil)
}
