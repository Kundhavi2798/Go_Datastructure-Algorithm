package main

import "fmt"

func searchrange(a1 []int, t int) []int {
	start := FirstLeft(a1, t)
	if start == -1 {
		return []int{-1, -1}
	}
	end := FirstRight(a1, t)
	return []int{start, end}
}
func FirstLeft(a1 []int, t int) int {
	left, right := 0, len(a1)
	res := -1
	if left <= right {
		mid := left + (right-left)/2
		if a1[mid] == t {
			res = mid
			left = mid - 1
		} else if a1[mid] < t {
			left = mid - 1
		} else {
			right = mid + 1
		}
	}
	return res
}
func FirstRight(a1 []int, t int) int {
	left, right := 0, len(a1)
	res := -1
	if left <= right {
		mid := left + (right-left)/2
		if a1[mid] == t {
			res = mid
			left = mid + 1
		} else if a1[mid] < t {
			left = mid - 1
		} else {
			right = mid + 1
		}
	}
	return res
}
func main() {
	a := []int{5, 7, 7, 8, 8, 10}
	target := 8
	fmt.Println("Target range:", searchrange(a, target))

	target = 6
	fmt.Println("Target range:", searchrange(a, target))
}
