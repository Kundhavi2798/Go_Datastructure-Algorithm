package main

import "fmt"

func main() {
	str := []string{"i", "am", "attending", "interview"}
	start, end := 0, len(str)-1
	for start < end {
		str[start], str[end] = str[end], str[start]
		start++
		end--
	}
	fmt.Println(str)
}
