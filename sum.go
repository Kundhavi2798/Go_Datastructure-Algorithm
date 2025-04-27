package main

import "fmt"

func main() {
	a := []int{1, 2, 4, 2, 4, 6}
	target := 10

	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			sum := a[i] + a[j]
			if sum == target {
				fmt.Printf("The idex which is add 1:%d 2:%d", i, j)
				break
			}
		}
	}
}
