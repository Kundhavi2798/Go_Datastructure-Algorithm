package main

import "fmt"

func reverseindexbased(input []int, k int) []int {
	for i := 0; i < len(input); i += k {
		left := i
		right := i + k - 1
		if right >= len(input) {
			right = len(input) - 1
		}
		for left < right {
			input[left], input[right] = input[right], input[left]
			left++
			right--
		}
	}

	return input
}
func topmulti(input []int) int {
	n := len(input)
	res := 1
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if input[j] < input[j+1] {
				input[j], input[j+1] = input[j+1], input[j]
			}
		}
	}
	fmt.Println(input)
	for i := 0; i < 3; i++ {
		res *= input[i]
	}
	return res
}
func flatten[T comparable](nested [][]T) []T {
	var data []T
	for _, i := range nested {
		data = append(data, i...)
	}
	return data
}
func main() {
	// num := []int{1, 4, 3, 2, 6, 5}
	// for n := len(num) - 1; n >= 0; n-- {
	// 	fmt.Print(num[n])
	// }
	Input := []int{10, 3, 5, 6, 20}
	// k := 3
	//val := reverseindexbased(Input, k)
	//fmt.Print(val)
	multi := topmulti(Input)
	fmt.Println(multi)
	nested := [][]int{
		{1, 2, 3},
		{4, 5},
		{6, 7, 8, 9},
	}

	flat := flatten(nested)
	fmt.Println(flat) 
}
