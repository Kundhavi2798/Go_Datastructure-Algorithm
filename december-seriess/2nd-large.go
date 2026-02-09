package main

import "fmt"

func main() {
	num := []int{12, 35, 1, 10, 34, 1}
	for i := 0; i < len(num)-1; i++ {
		for j := 0; j < len(num)-i-1; j++ {
			if num[j] > num[j+1] {
				num[j], num[j+1] = num[j+1], num[j]
			}
		}
	}
	count := 1
	large := num[len(num)-1]
	for i := len(num) - 2; i >= 0; i-- {
		if num[i] != large {
			count++
			if count == 3 {
				large = num[i]
			}
		}
	}
	fmt.Println("large2", large)
}
