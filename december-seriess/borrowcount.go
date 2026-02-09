package main

import "fmt"

func borrowCount(n1, n2 int) string {
	if n1 > n2 {
		return "Not Possible"
	}

	count := 0
	borrow := 0

	for n1 > 0 || n2 > 0 {
		d1 := n1 % 10
		d2 := n2 % 10

		d2 = d2 - borrow

		if d2 < d1 {
			count++
			borrow = 1
		} else {
			borrow = 0
		}

		n1 /= 10
		n2 /= 10
	}

	return fmt.Sprintf("%d", count)
}

func main() {
	fmt.Println(borrowCount(754, 897))  // 0
	fmt.Println(borrowCount(123, 1000)) // 3
	fmt.Println(borrowCount(95, 100))   // 2
	fmt.Println(borrowCount(500, 200))  // Not Possible
}
