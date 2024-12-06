package main

import "fmt"

func ispalinnum(s string) string {
	if len(s) == 0 {
		return ""
	}
	start := 0
	maxlen := 1
	for i := 0; i < len(s); i++ {
		left, right := i, i
		for left >= 0 && right < len(s) && s[left] == s[right] {
			if right-left+1 > maxlen {
				start = left
				maxlen = right - left + 1
			}
			left--
			right++
		}
		left, right = i, i+1
		for left >= 0 && right < len(s) && s[left] == s[right] {
			if right-left+1 > maxlen {
				start = left
				maxlen = right - left + 1
			}
			left--
			right++
		}
	}
	return s[start : start+maxlen]
}

func main() {
	s := "39878"
	fmt.Println("palindrome", ispalinnum(s))
}
