package main

import "fmt"

func main() {
	str := "pwwkew"
	charcheck := make(map[rune]bool)
	start, maxlen := 0, 0
	for i := 0; i < len(str); i++ {
		for charcheck[rune(str[i])] {
			delete(charcheck, rune(str[start]))
			start++
		}
		charcheck[rune(str[i])] = true
		currentLen := i - start + 1
		if currentLen > maxlen {
			maxlen = currentLen
		}
	}
	fmt.Println(maxlen)
}
