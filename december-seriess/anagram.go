package main

import "fmt"

func isAnagram(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	count := make(map[rune]int)
	for _, v := range s1 {
		count[v]++
	}
	for _, v := range s2 {
		count[v]--
		if count[v] < 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isAnagram("listen", "silent"))
	fmt.Println(isAnagram("rat", "car"))
}
