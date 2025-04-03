package main

import "fmt"

func sortnum(a []int) {
	n := len(a)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if a[j] < a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	fmt.Println(a)
	valProduct := 1
	for i := 0; i < 3; i++ {
		valProduct = valProduct * a[i]
	}
	fmt.Println(valProduct)
}
func main() {
	arr := []int{10, 3, 5, 6, 20}
	sortnum(arr)
	arr2 := []int{-10, -3, -5, -6, -20}
	sortnum(arr2)
	arr3 := []int{1, -4, 3, -6, 7, 0}
	sortnum(arr3)
}
