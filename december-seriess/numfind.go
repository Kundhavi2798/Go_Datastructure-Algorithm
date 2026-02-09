package main

import (
	"fmt"
	"strings"
)

func myAtoi(s string) {
	rem := strings.Trim(s, " ")
	fmt.Print(rem)
	// if rem >= 0 {
	// 	// fmt.Print(rem)
	// }

}

func main() {
	// myAtoi("42")
	myAtoi("  -042")
	// myAtoi("1337c0d3")
}
