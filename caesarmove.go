package main

import (
	"fmt"
)

func caesarCipher(s string, k int) string {
	encrypted := make([]rune, len(s))

	// Reduce unnecessary large shifts
	k = k % 26
	fmt.Println(k)

	for i, ch := range s {
		if ch >= 'a' && ch <= 'z' {
			ch = 'a' + (ch-'a'+rune(k))%26
		} else if ch >= 'A' && ch <= 'Z' {
			ch = 'A' + (ch-'A'+rune(k))%26
		}
		encrypted[i] = ch
	}

	return string(encrypted)
}

func main() {
	input := "middle-Outz"
	k := 2 // Shift by 2 positions
	output := caesarCipher(input, k)

	fmt.Println(output) // Expected Output: "okffng-Qwvb"
}
