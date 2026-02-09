package main

import "fmt"

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	// a, b := 0, 1
	// for i := 2; i < 10; i++ {
	// 	fmt.Print(a, " ")
	// 	a, b = b, a+b
	// }
	n := 10
	fmt.Print(fib(n))
}
