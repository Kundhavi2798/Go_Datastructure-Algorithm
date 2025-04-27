package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Name string
	Age  int
}

var users = []User{
	{"kundhavi", 20},
	{"sathi", 24},
}

func addUser(c *gin.Context) {
	var newUser User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	users = append(users, newUser)
	c.JSON(http.StatusOK, gin.H{
		"message": "User Added Successfully",
		"user":    users,
	})
}

func main() {
	router := gin.Default()
	router.GET("/user", func(context *gin.Context) {
		context.JSON(http.StatusOK, users)
	})

	router.POST("/user/add", addUser)
	router.Run(":8081")
}
