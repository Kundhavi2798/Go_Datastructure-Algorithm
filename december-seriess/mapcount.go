package main

import (
	"fmt"
)

func main() {
	a := []int{2, 6, 5, 8, 11, 2, 2, 3, 6, 6, 8, 8, 8, 8}
	seen := make(map[int]int)
	for _, v := range a {
		seen[v]++
	}
	fmt.Print(seen)
	max := 0
	result := 0
	for index, value := range seen {
		if value > max {
			max = value
			result = index
		}
	}
	fmt.Print(max, result)

}
