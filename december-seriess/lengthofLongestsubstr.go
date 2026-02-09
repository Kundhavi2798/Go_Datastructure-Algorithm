package main

import "fmt"

func lenofLongStr(s string) int {
	seen := make(map[byte]bool)
	start := 0
	maxLen := 0
	//already seen delete part
	for end := 0; end < len(s); end++ {
		for seen[s[end]] {
			delete(seen, s[start])
			start++
		}
		//seen set
		seen[s[end]] = true

		currentLen := end - start + 1
		if currentLen > maxLen {
			maxLen = currentLen
		}

	}
	return maxLen
}

func longestPalindrome(s string) string {
	seen := make(map[byte]bool)
	start := 0
	val := []int{}
	for end := 0; end < len(s); end++ {
		for seen[s[end]] {
			val = append(val, s[start])
			start++
		}
		seen[s[end]] = true
	}
			return val

}

func main() {
	// fmt.Println(lenofLongStr("abcabcab"))
	// fmt.Println(lenofLongStr("bbbbb"))
	// fmt.Println(lenofLongStr("pwwkew"))
	fmt.Println(longestPalindrome("babad"))

}
