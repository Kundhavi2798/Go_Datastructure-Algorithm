package main

import "fmt"

func isValid(s string) bool {
	stack := []rune{}
	bracketMap := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, ch := range s {
		fmt.Println(string(bracketMap[ch]))
		// If it's a closing bracket
		if open, ok := bracketMap[ch]; ok {
			fmt.Println(ok)
			fmt.Println(len(stack))
			if len(stack) == 0 || stack[len(stack)-1] != open {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			fmt.Print(string(ch))
			stack = append(stack, ch)
		}
	}
	return len(stack) == 0
}
func main() {
	s := "(]"
	fmt.Print(isValid(s))
}
