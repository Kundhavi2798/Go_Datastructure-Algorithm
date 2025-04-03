package main

import "fmt"

func main() {
	count := 0
	n := 10
	num := 2
	a := []int{}
	for count < n {
		isPrime := true
		for i := 2; i*i <= num; i++ {
			if num%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			a = append(a, num)
			count++
		}
		num++
	}
	fmt.Println(a)
}
