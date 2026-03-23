// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strings"
)

func commonPrefix(str []string) string {
	if len(str) == 0 {
		return ""
	}
	prefix := str[0]
	for i := 1; i < len(str); i++ {
		for !strings.HasPrefix(str[i], prefix) {
			if prefix == "" {
				return "Not found"
			}
			prefix = prefix[:len(prefix)-1]
		}

	}
	return prefix
}
func commonSuffix(str []string) string {
	if len(str) == 0 {
		return ""
	}
	sufix := str[0]
	for i := 1; i < len(str); i++ {
		for !strings.HasSuffix(str[i], sufix) {
			if sufix == "" {
				return "Not found"
			}
			sufix = sufix[1:]
		}

	}
	return sufix
}

func main() {
	str := []string{"kundhavi", "kundavai", "kundha"}

	prefix := commonPrefix(str)
	suffix := commonSuffix(str)

	fmt.Println("Prefix:", prefix)
	fmt.Println("Suffix:", suffix)
}
