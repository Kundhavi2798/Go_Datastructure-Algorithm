package main

import (
	"fmt"
	"strconv"
)

func superDigit(str string, k int) int {
	if len(str) == 1 {
		digit, _ := strconv.Atoi(str)
		return digit
	}
	sum := 0
	for _, d := range str {
		digit, _ := strconv.Atoi(string(d))
		sum += digit
	}
	return superDigit(strconv.Itoa(sum*k), 1)
}
func main() {
	str := "9875"
	k := 4
	fmt.Println(superDigit(str, k))
}
