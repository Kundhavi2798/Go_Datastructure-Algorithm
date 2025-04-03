package main

import (
	"fmt"
)

func minimumBribes(q []int) {
	n := len(q)
	total := 0
	sortedQ := make([]int, n)
	for i := 0; i < n; i++ {
		sortedQ[i] = i + 1
	}

	for i := 0; i < n; i++ {
		if q[i] == sortedQ[i] {
			continue
		} else if i+1 < n && q[i] == sortedQ[i+1] {
			sortedQ[i], sortedQ[i+1] = sortedQ[i+1], sortedQ[i]
			total++
		} else if i+2 < n && q[i] == sortedQ[i+2] {
			sortedQ[i+1], sortedQ[i+2] = sortedQ[i+2], sortedQ[i+1]
			sortedQ[i], sortedQ[i+1] = sortedQ[i+1], sortedQ[i]
			total += 2
		} else {
			fmt.Println("Too chaotic")
			return
		}
	}
	fmt.Println(total)
}

func main() {
	q := []int{2, 1, 5, 3, 4} // Example input
	minimumBribes(q)          // Expected output: 3
}
