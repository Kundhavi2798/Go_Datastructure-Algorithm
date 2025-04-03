package main

import "fmt"

func main() {
	s := "geeksforgeeks"
	str := []rune(s)
	fmt.Println(str)
	repstr := make(map[rune]int)
	for _, str := range s {
		repstr[str]++
	}
	fmt.Println(repstr)
	for _, str := range s {
		if repstr[str] == 1 {
			fmt.Println(string(str))
			break
		}
	}

}
