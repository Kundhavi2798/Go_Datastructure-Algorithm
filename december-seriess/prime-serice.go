package main

import "fmt"

func main() {
	count := 0
	num := 2
	ser := 10
	val := []int{}
	for count < ser {
		isprime := true
		for i := 2; i*i <= num; i++ {
			if num%i == 0 {
				isprime = false
				break
			}
		}
		if isprime {
			val = append(val, num)
			count++
		}
		num++
	}
	fmt.Println(val)
}
