package main

import "fmt"

func permute(s []rune, l int, r int) {
	if l == r {
		fmt.Println(string(s))
		return
	}
	for i := l; i <= r; i++ {
		s[l], s[i] = s[i], s[l]
		permute(s, l+1, r)
		s[l], s[i] = s[i], s[l]
	}

}
func main() {
	str := "ABC"
	permute([]rune(str), 0, len(str)-1)
}
