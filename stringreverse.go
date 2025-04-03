package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "..geeks..for.geeks."
	trimesplce := strings.Trim(str, ".")
	fmt.Println(trimesplce)
	splitStr := strings.Split(trimesplce, ".")
	fmt.Println(splitStr)
	var removeSpace []string
	for _, splitStr := range splitStr {
		if splitStr != "" {
			removeSpace = append(removeSpace, splitStr)
		}
	}
	n := len(removeSpace)
	for i := 0; i < n/2; i++ {
		removeSpace[i], removeSpace[n-i-1] = removeSpace[n-i-1], removeSpace[i]
	}
	reverstr := strings.Join(removeSpace, ".")
	fmt.Println(reverstr)
}
