package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func hellohandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello word")

}
func helloword(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Helloo kundhavi",
		"Name":    "Kundhavi",
		"Role":    "SpftwareDeveloper",
		"Age":     27,
	})
}
func main() {
	//Mux handler
	// r := mux.NewRouter()
	// r.HandleFunc("/hello", hellohandler).Methods("GET")
	// fmt.Println("server listening port http://localhost:8080")
	// http.ListenAndServe(":8080", r)
	r := gin.Default()
	r.GET("/hello", helloword)
	r.Run(":8080")
}
