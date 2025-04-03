package main

import (
	"fmt"
	"math"
)

func main() {
	num := 153
	numdig := len(fmt.Sprintf("%d", num))
	sum, temp := 0, num
	for temp != 0 {
		digit := temp % 10
		sum += int(math.Pow(float64(digit), float64(numdig)))
		temp /= 10
	}
	if sum == num {
		fmt.Println("Its an anagram")
	} else {
		fmt.Println("Its not an anagram")
	}
}
