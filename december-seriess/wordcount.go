package main

// import "fmt"

// func main() {
// 	a := "aabbbbeeeeffggg"
// 	count := 1
// 	res := ""
// 	for i := 1; i < len(a); i++ {
// 		if a[i] == a[i-1] {
// 			count++
// 		} else {
// 			res += string(a[i-1])
// 			res += fmt.Sprintf("%d", count)
// 			count = 1
// 		}
// 	}
// 	res += string(a[len(a)-1])
// 	res += fmt.Sprintf("%d", count)
// 	fmt.Println(res)
// }

// package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func main() {
	s := "a1b4e4f2g3"
	res := ""
	count := ""

	for _, ch := range s {
		if unicode.IsDigit(ch) {
			count += string(ch)
		} else {
			if count != "" {
				n, _ := strconv.Atoi(count)
				for i := 0; i < n; i++ {
					res += string(ch)
				}
				count = ""
			}
		}
	}

	fmt.Println(res)
}
