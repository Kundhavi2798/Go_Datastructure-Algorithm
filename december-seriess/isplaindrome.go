package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	original := x
	rev := 0
	for x != 0 {
		digit := x % 10
		rev = rev*10 + digit
		x = x / 10
	}
	return original == rev
}
func main() {
	fmt.Println(isPalindrome(121))  // true
	fmt.Println(isPalindrome(123))  // false
	fmt.Println(isPalindrome(1331)) // true
	fmt.Println(isPalindrome(-121))
}
