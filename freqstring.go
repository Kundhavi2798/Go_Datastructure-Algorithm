package main

import "fmt"

func freqcountStr(str string) string {
	found := make(map[rune]int)
	strcov := []rune(str)
	res := ""
	fmt.Println()
	for _, num := range strcov {
		found[num]++
	}
	for _, num := range strcov {
		if found[num] == 1 {
			res = string(num)
			break
		} else if found[num] > 1 {
			res = "No Repeating"
		}
	}
	return res
}
func main() {
	a1 := "swiss"
	a2 := "aabcc"
	fmt.Println("a1", freqcountStr(a1))
	fmt.Println("a2", freqcountStr(a2))
}
