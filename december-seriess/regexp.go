package main

import (
	"fmt"
	"regexp"
)

func regexpCheck(s, p string) bool {
	ok, _ := regexp.MatchString("^"+p+"$", s)
	return ok
}

func main() {
	fmt.Println(regexpCheck("aa", "a"))  //false
	fmt.Println(regexpCheck("aa", "*a")) //true
	fmt.Println(regexpCheck("aa", ".*")) //true
}
