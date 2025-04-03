package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file := "sample.txt"
	fileData, err := os.Create(file)
	if err != nil {
		fmt.Println("cant open file", err)
	}
	writeData, err := fileData.WriteString("Wellcome")
	fmt.Println(writeData)
	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("cant open file", err)
	}
	fmt.Println(string(data))
}
