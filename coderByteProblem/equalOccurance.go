package main

import "fmt"

func main() {
	str := "abacbc"
	freq := make(map[rune]int)
	conData := []rune(str)
	for _, i := range conData {
		freq[i]++
	}
	val := []int{}

	for _, j := range freq {
		val = append(val, j)
	}
	allcheck := true
	fmt.Println(val)
	for i := 1; i < len(val); i++ {
		if val[i] != val[0] {
			allcheck = false
			break
		}
	}
	fmt.Println(allcheck)
}
