package main

import (
	"fmt"
)

func flippingMatrix(matrix [][]int) int {
	n := len(matrix) / 2 // Half the matrix size
	maxSum := 0

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			// Take the maximum value among the four symmetric positions
			maxSum += max(
				matrix[i][j],
				matrix[i][2*n-j-1],
				matrix[2*n-i-1][j],
				matrix[2*n-i-1][2*n-j-1],
			)
		}
	}

	return maxSum
}

// Utility function to find max of four numbers
func max(a, b, c, d int) int {
	m := a
	if b > m {
		m = b
	}
	if c > m {
		m = c
	}
	if d > m {
		m = d
	}
	return m
}

func main() {
	// Example test case
	matrix := [][]int{
		{112, 42, 83, 119},
		{56, 125, 56, 49},
		{15, 78, 101, 43},
		{62, 98, 114, 108},
	}

	fmt.Println(flippingMatrix(matrix)) // Output: 414
}
