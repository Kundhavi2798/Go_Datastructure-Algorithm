package main

import (
	"fmt"
	"math"
)

func MinWindowSubstring(strArr []string) string {
	N := strArr[0]
	K := strArr[1]

	// Frequency map for characters in K
	neededChars := make(map[rune]int)
	for _, char := range K {
		neededChars[char]++
	}

	// Variables for the sliding window
	left := 0
	minLength := math.MaxInt32
	minSubstring := ""
	charCount := 0

	// Frequency map for the current window
	windowChars := make(map[rune]int)

	// Expand the window by moving the right pointer
	for right, char := range N {
		windowChars[char]++
		if neededChars[char] > 0 && windowChars[char] <= neededChars[char] {
			charCount++
		}

		// When all characters from K are included in the window
		for charCount == len(K) {
			// Update the smallest substring
			if right-left+1 < minLength {
				minLength = right - left + 1
				minSubstring = N[left : right+1]
			}

			// Shrink the window from the left
			leftChar := rune(N[left])
			windowChars[leftChar]--
			if neededChars[leftChar] > 0 && windowChars[leftChar] < neededChars[leftChar] {
				charCount--
			}
			left++
		}
	}

	return minSubstring
}

func main() {
	strArr1 := []string{"aaabaaddae", "aed"}
	fmt.Println(MinWindowSubstring(strArr1)) // Output: "dae"

	strArr2 := []string{"aabdccdbcacd", "aad"}
	fmt.Println(MinWindowSubstring(strArr2)) // Output: "aabd"
}
