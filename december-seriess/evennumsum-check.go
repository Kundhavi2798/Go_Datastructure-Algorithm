package main

import (
	"fmt"
)

// Input: num = 4
// Output: 2
// Explanation:
// The only integers less than or equal to 4 whose digit sums are even are 2 and 4.
func main() {
	num := 10
	rem := num % 2
	if rem == 0 {
		num1 := num / 2
		sum := num1 + num1
		if sum == num {
			fmt.Println(num1)
		} else {
			fmt.Print("Not equal to number")
		}
	} else {
		fmt.Println("not a even number")
	}

}
