package main

import (
	"fmt"
	"sort"
)

func gridChallenge(grid []string) string {
	n := len(grid)
	sortedGrid := make([][]rune, n)

	// Step 1: Sort each row
	for i := 0; i < n; i++ {
		row := []rune(grid[i])
		sort.Slice(row, func(a, b int) bool {
			return row[a] < row[b]
		})
		sortedGrid[i] = row
	}

	// Step 2: Check if each column is sorted
	for col := 0; col < len(sortedGrid[0]); col++ { // ✅ Fixed condition (col < len(sortedGrid[0]))
		for row := 1; row < n; row++ {
			if sortedGrid[row][col] < sortedGrid[row-1][col] {
				return "NO"
			}
		}
	}

	return "YES"
}

// Test cases
func main() {
	testCases := [][]string{
		{
			"ebacd",
			"fghij",
			"olmkn",
			"trpqs",
			"xywuv",
		}, // Expected Output: YES
		{
			"mpxz",
			"abcd",
			"wlmf",
		}, // Expected Output: NO
	}

	for _, grid := range testCases {
		fmt.Println(gridChallenge(grid))
	}
}
