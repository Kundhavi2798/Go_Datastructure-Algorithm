package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := 1230
	numStr := strconv.Itoa(num)
	numtoStr := map[rune]string{
		'0': "Zero",
		'1': "One",
		'2': "Two",
		'3': "Three",
		'4': "Four",
		'5': "Five",
		'6': "Six",
		'7': "Seven",
		'8': "Eight",
		'9': "Nine",
	}
	for _, ch := range numStr {
		//if word, exists := numtoStr[ch]; exists {
		//	fmt.Print(word, " ")
		//}
		fmt.Println(numtoStr[ch])
	}
}
