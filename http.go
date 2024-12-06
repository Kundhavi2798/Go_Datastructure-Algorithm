package main

import "fmt"

// Helper function to check if a cell is within bounds
func isValid(x, y, rows, cols int) bool {
	return x >= 0 && x < rows && y >= 0 && y < cols
}

// DFS to mark all 'O' cells connected to the boundary
func dfs(grid [][]byte, x, y, rows, cols int) {
	if !isValid(x, y, rows, cols) || grid[x][y] != 'O' {
		return
	}

	// Mark the cell as visited by changing it to a temporary value
	grid[x][y] = 'T'

	// Explore all 4 directions
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for _, dir := range directions {
		dfs(grid, x+dir[0], y+dir[1], rows, cols)
	}
}

// Function to replace 'O' islands with 'X'
func replaceIslands(grid [][]byte) [][]byte {
	rows := len(grid)
	cols := len(grid[0])

	// Mark all boundary-connected 'O's as 'T'
	for i := 0; i < rows; i++ {
		if grid[i][0] == 'O' {
			dfs(grid, i, 0, rows, cols)
		}
		if grid[i][cols-1] == 'O' {
			dfs(grid, i, cols-1, rows, cols)
		}
	}
	for j := 0; j < cols; j++ {
		if grid[0][j] == 'O' {
			dfs(grid, 0, j, rows, cols)
		}
		if grid[rows-1][j] == 'O' {
			dfs(grid, rows-1, j, rows, cols)
		}
	}

	// Replace all 'O's with 'X' and revert 'T' to 'O'
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 'O' {
				grid[i][j] = 'X'
			} else if grid[i][j] == 'T' {
				grid[i][j] = 'O'
			}
		}
	}

	return grid
}

func main() {
	// Example Input
	grid := [][]byte{
		{'X', 'X', 'O', 'O'},
		{'X', 'O', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
		{'X', 'X', 'X', 'X'},
	}

	// Replace Islands
	result := replaceIslands(grid)

	// Print the Result
	for _, row := range result {
		for _, cell := range row {
			fmt.Print(string(cell), " ")
		}
		fmt.Println()
	}
}
