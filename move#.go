package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Move#Hash#to#Front"
	count := strings.Count(str, "#")
	fmt.Println(count)
	removed := strings.ReplaceAll(str, "#", "")
	fmt.Println(removed)
	recomputedstr := strings.Repeat("#", count)
	fmt.Println(recomputedstr + removed)
}
