package main

import (
	"fmt"
	"sort"
)

func prefix(s []string) string {
	left, right := s[0], s[len(s)-1]
	i := 0
	for i < len(left) && i < len(right) && left[i] == right[i] {
		i++
	}
	return left[:i]
}

func sufix(s []string) string {
	left, right := s[len(s)-1], s[0]
	fmt.Println(left, right)
	i := len(left) - 1
	j := len(right) - 1
	for i >= 0 && j >= 0 && left[i] == right[j] {
		i--
		j--
	}
	fmt.Println(i)
	return right[i+1:]
}
func main() {
	str := []string{"flowerrrr", "flowrsss", "flightrrrrrr"}
	sort.Strings(str)
	fmt.Println(prefix(str))
	fmt.Println(sufix(str))
}
