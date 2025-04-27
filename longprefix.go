package main

import (
	"fmt"
	"sort"
)

func prefixfoound(str []string) string {
	if len(str) == 0 {
		fmt.Println("no prefix form")
	}
	sort.Strings(str)
	fmt.Println(str)
	left, right := str[0], str[len(str)-1]
	fmt.Println(left, right)
	i := 0
	for i < len(left) && i < len(right) && left[i] == right[i] {
		i++
	}
	return left[:i]
}
func main() {
	strArr := []string{"geeksforgeeks", "geeks", "geek", "geezer"}
	fmt.Println(prefixfoound(strArr))
}
