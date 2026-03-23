package main

import (
	"fmt"
)

func main() {
	strs := []string{"flowerrrr", "flowrrr", "flightrrrr"}
	start := strs[0]
	suffix := ""
	for i := 0; i < len(start); i++ {
		for j := 1; j < len(strs); j++ {
			if strs[j][len(strs[j])-1-i] != start[len(start)-1-i] {
				fmt.Print(suffix)
				return
			}
		}
		suffix = string(start[len(start)-1-i]) + suffix
	}
}
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	first := strs[0]

	for i := 0; i < len(first); i++ {

		for j := 1; j < len(strs); j++ {

			if i >= len(strs[j]) || strs[j][i] != first[i] {
				return first[:i]
			}
		}
	}

	return first
}

func mergeAlternately(word1 string, word2 string) string {
	res := ""
	i := 0
	for i < len(word1) || i < len(word2) {
		if i < len(word1) {
			res += string(word1[i])
		}
		if i < len(word2) {
			res += string(word2[i])
		}
		i++
	}
	return res
}
