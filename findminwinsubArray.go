package main

import (
	"fmt"
	"strings"
)

func compareStrings(str1, str2 string) bool {
	for _, v := range str1 {
		if !strings.Contains(str2, string(v)) {
			return false
		} else {
			str2 = strings.Replace(str2, string(v), "", 1)
		}
	}
	return true
}

func MinWindowSubstringcheck(strArr []string) string {
	for minLen := len(strArr[1]); minLen <= len(strArr[0]); minLen++ { // changed `<` to `<=` to include full length cases
		for i := 0; i <= len(strArr[0])-minLen; i++ {
			inputFormat := strArr[0][i : i+minLen]
			if compareStrings(strArr[1], inputFormat) {
				return inputFormat
			}
		}
	}
	return ""
}

func main() {
	fmt.Println(MinWindowSubstringcheck([]string{"aaabaaddae", "aed"}))
	fmt.Println(MinWindowSubstringcheck([]string{"aabdccdbcacd", "aad"})) // fixed closing parenthesis
}
