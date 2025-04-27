package main

import "fmt"

func main() {
	arr := []int{5, 3, 6, 7, 5, 3}
	n := len(arr)
	sum := 0
	for i := 0; i < n; i++ {
		sum += arr[i]
	}
	total := float64(sum) / float64(n)
	fmt.Println(total)

}
