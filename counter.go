package main

import (
	"fmt"
	"regexp"
	"unicode"
)

func isvalidpass(pass string) bool {
	fmt.Println("passowrd", len(pass))
	if len(pass) < 8 {
		return false
	}
	if !unicode.IsUpper(rune(pass[0])) {
		return false
	}
	hasnum := regexp.MustCompile(`[0-9]`).MatchString(pass)
	if !hasnum {
		return false
	}
	hasspe := regexp.MustCompile(`[%&@*$#]`).MatchString(pass)
	if !hasspe {
		return false
	}
	return true
}

func main() {
	//a := 67811

	//num := a % 100
	//num2 := num%10 + num/10
	//if num2 == 3 || num2 == 8 {
	//	fmt.Println("Lucky Draw Winner")
	//} else {
	//	fmt.Println("Not a Lucky Draw Winner")
	//}
	//var n int
	//fmt.Print("enter the num:")
	//fmt.Scan(&n)
	//a := make([]int, n)
	//fmt.Print("enter the array value :")
	//for i := 0; i < n; i++ {
	//	fmt.Scan(&a[i])
	//}
	//fmt.Println("the array :", a)
	//for i := 0; i < n-1; i++ {
	//	for j := 0; j < n-1-i; j++ {
	//		if a[j] < a[j+1] {
	//			a[j], a[j+1] = a[j+1], a[j]
	//		}
	//	}
	//}
	//fmt.Println("the array after sort :", a)
	//length := len(a)
	//fmt.Println("center number", a[length/2])
	var password string
	fmt.Print("Enter password:")
	fmt.Scan(&password)
	if isvalidpass(password) {
		fmt.Println("valid")
	} else {
		fmt.Println("not valid")
	}
}
