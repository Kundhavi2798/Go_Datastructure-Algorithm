package main

import "fmt"

func main() {
	a, b := 0, 1
	arr := []int{}
	for i := 2; i < 10; i++ {
		a, b = b, a+b
		arr = append(arr, a)
	}
	fmt.Println(arr)
}
