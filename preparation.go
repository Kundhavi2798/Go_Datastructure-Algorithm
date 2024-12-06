package main

import "fmt"

func searchelement(s []int, t int) []int {
	start := firstpart(s, t)
	if start == -1 {
		return []int{-1, -1}
	}
	end := secondpart(s, t)
	return []int{start, end}
}
func firstpart(s []int, t int) int {
	low, high := 0, len(s)
	res := -1
	fmt.Println("low and high", low, high)

	if low <= high {
		mid := low + (high-low)/2
		fmt.Println("mid value", mid)

		if s[mid] == t {
			res = mid
			high = mid - 1
		} else if s[mid] < t {
			low = mid - 1
		} else {
			high = mid + 1
		}
	}
	fmt.Println(res)
	return res
}
func secondpart(s []int, t int) int {
	low, high := 0, len(s)
	res := -1
	fmt.Println("low and high", low, high)
	if low <= high {
		mid := low + (high-low)/2
		fmt.Println("mid value", mid)
		if s[mid] == t {
			res = mid
			low = mid + 1
		} else if s[mid] < t {
			low = mid - 1
		} else {
			high = mid + 1
		}
	}
	fmt.Println("res", res)
	return res
}
func checkpalin(s string) bool {
	n := len(s)
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-1-i] {
			return false
		}
	}
	return true
}
func searchnum(a []int, t int) bool {
	for i, val := range a {
		if val == t {
			a = append(a[:i], a[i+1:]...)
			fmt.Println(a)
		}
	}
	return false
}
func insertnum(a []int, val int, inx int) []int {
	a = append(a[:inx], append([]int{val}, a[inx:]...)...)
	return a
}
func main() {
	//arr := []string{"00900", "bsds", "aaa"}
	//for _, r := range arr {
	//	if checkpalin(r) {
	//		fmt.Println("palin")
	//	} else {
	//		fmt.Println("not palin")
	//	}
	//}
	s := []int{17, 15, 0, 0, 8, 9, 12}
	n := len(s)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if s[j] < s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
	fmt.Println("array", s)
}
