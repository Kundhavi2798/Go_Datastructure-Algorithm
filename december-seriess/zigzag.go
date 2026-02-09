package main

import "fmt"

func main() {
	s := "PAYPALISHIRING"
	numRows := 3
	rows := make([]string, numRows)
	currentrow := 0
	down := false

	for _, ch := range s {
		rows[currentrow] += string(ch)
		if currentrow == 0 || currentrow == numRows-1 {
			down = !down
		}
		if down {
			currentrow++
		} else {
			currentrow--
		}
	}
	res := ""
	for _, s := range rows {
		res += s
	}
	fmt.Print(res)
}
