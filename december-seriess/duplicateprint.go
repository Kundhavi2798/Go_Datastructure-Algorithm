package main

import (
	"fmt"
)

func main() {
	str := "kundhavikv"
	convstr := []rune(str)

	seen := make(map[rune]int)
	for _, ch := range convstr {
		seen[ch]++
	}
	fmt.Println(seen)
	max := 0
	res := []string{}
	for key, val := range seen {
		fmt.Println(string(key), val)
		if val > max {
			max = val
			res = []string{string(key)}

		} else if val == max {
			res = append(res, string(key))
		}

	}
	fmt.Print(res)
}
