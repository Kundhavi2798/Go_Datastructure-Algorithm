package main

import (
	"fmt"
	"strings"
)

func rotate(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	return strings.Contains(s1+s1, s2)
}
func main() {
	s1 := "abcd"
	s2 := "cdab"
	fmt.Println(rotate(s1, s2))
	s3, s4 := "aab", "aba"
	fmt.Println(rotate(s3, s4))
	s5, s6 := "abcd", "acbd"
	fmt.Println(rotate(s5, s6))

}
