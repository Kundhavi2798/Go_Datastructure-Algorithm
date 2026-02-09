package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 4, 3, 4, 3, 4, 4, 5}
	k := 3
	for i := 0; i < k; i++ {
		for k := 0; k < len(a)-1; k++ {
			for j := 0; j < len(a)-k-1; j++ {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	fmt.Print(a)

}
