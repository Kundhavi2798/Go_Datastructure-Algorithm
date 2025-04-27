package main

import (
	"fmt"
	"sort"
)

func maxProduct(a []int) int {
	sort.Ints(a)
	product := a[len(a)-1] * a[len(a)-2]
	return product
}

func moveZero(a []int) []int {
	zero := 0
	for i := 0; i < len(a); i++ {
		if a[i] != 0 {
			a[i], a[zero] = a[zero], a[i]
			zero++
		}
	}
	return a
}

func main() {
	arr := []int{3, 5, -2, 8, 10, -7}
	fmt.Println("Max Product:", maxProduct(arr))
	arr2 := []int{0, 1, 0, 3, 12}
	fmt.Println("the final array:", moveZero(arr2))

}
