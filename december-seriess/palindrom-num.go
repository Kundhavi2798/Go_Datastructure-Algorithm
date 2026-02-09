package main

import "fmt"

func ispalinNum(n int) bool {
	original := n
	rev := 0
	for n != 0 {
		digit := n % 10
		rev = rev*10 + digit
		n = n / 10
	}
	return original == rev
}
func isPalinstr(s string) bool {
	val := []rune(s)
	for i := 0; i < len(s)-1; i++ {
		for j := 0; j < len(s)-i-1; j++ {
			val[j], val[j+1] = val[j+1], val[j]
		}
	}
	return string(val) == s
}
func main() {
	// n := 121
	// fmt.Println(ispalinNum(n))
	// str := "amma"
	// fmt.Println(isPalinstr(str))
	strarr := []string{"amma", "akka", "appa", "kasdasd"}
	for _, v := range strarr {
		fmt.Println(isPalinstr(v))
	}
}
