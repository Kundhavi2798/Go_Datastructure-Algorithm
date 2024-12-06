package main

import "fmt"

func main() {
	a := []int{0, 0, 0, 1, 1, 1, 2, 2, 3}
	count := make(map[int]int)
	for _, freq := range a {
		count[freq]++
	}
	fmt.Println(count)
	var result, maxFreq int
	for num, val := range count {
		if val > 1 {
			if val > maxFreq || (val == maxFreq && num > result) {
				maxFreq = val
				result = num
				fmt.Println("value", maxFreq, result)
			}
		}
	}
	fmt.Println(result)
}
