package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	str := "go is fun and go is powerful"
	spltstring := strings.Split(str, " ")
	fmt.Println(spltstring)
	sort.Slice(spltstring, func(i, j int) bool {
		return spltstring[i] > spltstring[j]
	})

	fmt.Println(spltstring)
}
