package main

import (
	"bytes"
	"fmt"
)

func main() {
	str1 := []byte{'I', 'N', 'T', 'E', 'R', 'V', 'I', 'E', 'W'}
	str2 := []byte{'B', 'I', 'T'}
	found := bytes.Contains(str1, str2)
	if found {
		fmt.Println("Present")
	} else {
		fmt.Println("Not Present")
	}
}
