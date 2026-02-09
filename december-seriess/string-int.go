package main

import (
	"fmt"
	"strings"
)

func findint(s string) int {
	sign := 1
	s = strings.TrimSpace(s)
	// val := 0
	if strings.HasPrefix(s, "-") {
		sign = -1
		s = s[1:]
	}
	num := 0
	for _, ch := range s {
		if ch < '0' || ch > '9' {
			break
		}
		num = num*10 + int(ch-'0')
	}

	return sign * num

}

func main() {
	s := "  -123213ssd12"
	fmt.Println(findint(s))
}
