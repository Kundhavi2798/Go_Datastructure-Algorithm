package main

import "fmt"

func isPlain(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
func extraCharsForPalindrome(s string) string {
	n := len(s)
	for i := 1; i < n; i++ {
		if isPlain(s[i:]) {
			extra := ""
			for j := i - 1; j >= 0; j-- {
				extra += string(s[j])
			}
			return extra
		}
	}
	return ""
}
func main() {
	fmt.Println(extraCharsForPalindrome("abede")) // ba
	fmt.Println(extraCharsForPalindrome("abcfe"))
}
