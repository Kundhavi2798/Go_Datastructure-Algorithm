package main

import "fmt"

func main() {
	//str := "abba"
	str := "abc"
	str_rune := []rune(str)
	for i := 0; i < len(str)-1; i++ {
		for j := 0; j < len(str)-i-1; j++ {
			str_rune[j], str_rune[j+1] = str_rune[j+1], str_rune[j]
		}
	}
	if str == string(str_rune) {
		fmt.Println("its palindrome")
	} else {
		fmt.Println("its not palindrome")

	}

}
