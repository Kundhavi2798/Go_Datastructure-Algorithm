package main

import (
	"encoding/json"
	"fmt"
)

type UserData struct {
	Id            int    `json:"ID"`
	Title         string `json:"Title"`
	Author        string `json:"Author"`
	PublishedYear int    `json:"PublishedYear"`
}

var JsonData = `[	{ID: 1, Title: "The Go Programming Language", Author: "Alan Donovan", PublishedYear: 2015},
	{ID: 2, Title: "Clean Code", Author: "Robert C. Martin", PublishedYear: 2008}]`

func main() {
	var userData []UserData
	err := json.Unmarshal([]byte(JsonData), userData)
	if err != nil {
		fmt.Println("errors", err)
	}
	for _, details := range userData {
		fmt.Println(details)
	}
}
