package main

import "fmt"

func romantoInt() {
	str := "MCMXCIV"
	roman := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	n := len(str)
	res := 0

	for i := 0; i < n; i++ {
		if i < n-1 && roman[str[i]] < roman[str[i+1]] {
			res -= roman[str[i]]
		} else {
			res += roman[str[i]]
		}
	}
	fmt.Print(res)
}
func InttoRoman() {
	num := 1994
	vals := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	syms := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	res := ""
	for i := 0; i < len(vals); i++ {
		for num >= vals[i] {
			num -= vals[i]
			res += syms[i]
		}
	}
	fmt.Print(res)
}
func main() {
	romantoInt()
	InttoRoman()
}
