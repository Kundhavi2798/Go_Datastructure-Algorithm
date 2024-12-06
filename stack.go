package main

import "fmt"

type Stack struct {
	val []int
}

func (s *Stack) push(n int) {
	s.val = append(s.val, n)
}
func (s *Stack) pop() int {
	if len(s.val) == 0 {
		fmt.Println("stack is empty")
	}
	top := len(s.val) - 1
	s.val = s.val[:len(s.val)-1]
	return top
}
func (s *Stack) peek() int {
	if len(s.val) == 0 {
		fmt.Println("stack is empty")
	}
	top := len(s.val) - 1
	return s.val[top]
}
func main() {
	stack := &Stack{}
	stack.push(10)
	stack.push(10)
	stack.push(10)
	fmt.Println("values", stack.val)
	stack.pop()
	fmt.Println("values", stack.val)
	stack.peek()
	fmt.Println("values", stack.val)

}
