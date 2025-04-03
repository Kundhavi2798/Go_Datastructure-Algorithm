package main

import (
	"fmt"
)

// Function to check if a string is a palindrome
func isPalindrome(s string) bool {
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

// Function to find the index to remove
func palindromeIndex(s string) int {
	if isPalindrome(s) {
		return -1 // Already a palindrome
	}

	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			// Check by removing left character
			if isPalindrome(s[:left] + s[left+1:]) {
				return left
			}
			// Check by removing right character
			if isPalindrome(s[:right] + s[right+1:]) {
				return right
			}
			return -1 // If neither works, return -1
		}
		left++
		right--
	}
	return -1
}

func main() {
	// Test cases
	fmt.Println(palindromeIndex("aaab"))    // Output: 3
	fmt.Println(palindromeIndex("baa"))     // Output: 0
	fmt.Println(palindromeIndex("racecar")) // Output: -1 (Already palindrome)
	fmt.Println(palindromeIndex("abca"))    // Output: 1
}
