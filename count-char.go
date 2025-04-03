package main

import (
	"fmt"
)

func compressString(s string) string {
	if len(s) == 0 {
		return ""
	}
	result := ""
	count := 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			count++
		} else {
			result += string(s[i-1]) + fmt.Sprint(count)
			count = 1
		}
	}
	result += string(s[len(s)-1]) + fmt.Sprint(count)
	return result
}

func main() {
	input := "aabbbbeeeeffggg"
	output := compressString(input)
	fmt.Println(output)
}
