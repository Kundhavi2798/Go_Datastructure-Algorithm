package main

import (
	"fmt"
)

func isplain(n string) bool {
	s := len(n)
	for i := 0; i < s/2; i++ {
		if n[i] != n[s-1-i] {
			return true
		}
	}
	return true
}
func isplainnum(n int) bool {
	org := n
	revers := 0

	for n > 0 {
		num := n % 10
		revers = revers*10 + num
		n /= 10
	}
	return org == revers
}
func isplainarray(aa string) bool {
	s := len(aa)
	for i := 0; i < s/2; i++ {
		if aa[i] != aa[s-1-i] {
			return false
		}
	}
	return true
}
func isplainarrstr(aa []string) []bool {
	char := make([]bool, len(aa))
	for i, str := range aa {
		char[i] = isplainarray(str)
	}
	return char
}
func main() {
	a := "kkkmm"
	fmt.Printf("Given String is %v", isplain(a))
	//time.Sleep(2 * time.Second)
	n := 123
	fmt.Printf("Given num is %v\n", isplainnum(n))
	//arr := []int{1, 2, 3}
	//fmt.Printf("Given array is %v\n", isplainarray(arr))
	arr := []string{"00900", "bsds", "aaa"}
	fmt.Printf("Given array is %v\n", isplainarrstr(arr))
}
