package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type jsonSampleData struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	url := "https://dummy-json.mock.beeceptor.com/todos"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("get error", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error status code:", resp.StatusCode)
		return
	}
	readData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("cant read", err)
		return
	}
	var jsonResponse []jsonSampleData
	errjson := json.Unmarshal(readData, &jsonResponse)
	if errjson != nil {
		fmt.Println("cant unmarshal", errjson)
		return
	}
	for _, listItem := range jsonResponse {
		fmt.Println(listItem)
		//fmt.Printf("userid:%d\n ID :%d \n Title :%s Completed %v", listItem.UserId, listItem.Id, listItem.Title, listItem.Completed)
	}
}
