package main

import (
	"fmt"
	"strings"
	"unicode"
)

func modifyString(s string) string {
	// Step 1: Reverse each word in the string
	words := strings.Fields(s)
	for i, word := range words {
		runes := []rune(word)
		for j, k := 0, len(runes)-1; j < k; j, k = j+1, k-1 {
			runes[j], runes[k] = runes[k], runes[j]
		}
		words[i] = string(runes)
	}
	reversedString := strings.Join(words, " ")

	// Step 2: Replace even index characters with uppercase and odd index characters with lowercase
	modified := []rune(reversedString)
	for i, char := range modified {
		if unicode.IsLetter(char) {
			if i%2 == 0 {
				modified[i] = unicode.ToUpper(char)
			} else {
				modified[i] = unicode.ToLower(char)
			}
		}
	}

	// Step 3 & 4: Replace spaces and other characters based on their index
	for i, char := range modified {
		if unicode.IsSpace(char) {
			if i%2 == 0 {
				modified[i] = '#'
			} else {
				modified[i] = '*'
			}
		} else if !unicode.IsLetter(char) {
			if i%2 == 0 {
				modified[i] = '@'
			} else {
				modified[i] = ':'
			}
		}
	}

	return string(modified)
}

func main() {
	var s string
	fmt.Println("Enter the string:")
	fmt.Scanln(&s)

	result := modifyString(s)
	fmt.Println("Modified string:", result)
}
