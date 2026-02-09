package main

import (
	"fmt"
)

func main() {
	a := []int{2, 1, 5, 1, 3, 2}
	k := 3
	max := 0
	for i := 0; i <= len(a)-k; i++ {
		sum := 0

		for j := i; j < i+k; j++ {
			sum += a[j]
		}
		if sum > max {
			max = sum
		}
	}
	fmt.Print(max)

}
