package main

import "fmt"

func main() {
	a := []int{34, -50, 42, 14, -5, 86}
	sliednum := 4
	length := len(a)
	if sliednum > length {
		fmt.Println("cant find maxsum the sliding is not ")
	}
	cursum := 0
	for i := 0; i < sliednum; i++ {
		cursum += a[i]
	}
	fmt.Println("sume", cursum)
	maxsum := cursum
	for i := sliednum; i < length; i++ {
		cursum += a[i] - a[i-sliednum]
		if cursum > maxsum {
			maxsum = cursum
		}
	}
	fmt.Println("maxsum", maxsum)
}
