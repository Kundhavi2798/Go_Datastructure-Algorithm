package main

import (
	"fmt"
	"time"
)

func checkpalinstr(s string) string {
	revstr := []rune(s)
	for i := 0; i < len(s)-1; i++ {
		for j := 0; j < len(s)-i-1; j++ {
			revstr[j], revstr[j+1] = revstr[j+1], revstr[j]
		}
	}
	if string(revstr) == s {
		return string(revstr)
	} else {
		return ""
	}
}

func checkpalindrome(s []string) {
	checkdata := make(chan []string)
	res := []string{}
	for i := 0; i < len(s); i++ {
		str := checkpalinstr(s[i])
		res = append(res, str)
	}
	go func() {
		checkdata <- res
	}()
	fmt.Println(<-checkdata)

}
func main() {
	str := []string{"racecar", "hello", "level", "world", "madam"}
	go checkpalindrome(str)
	time.Sleep(time.Second)
}
