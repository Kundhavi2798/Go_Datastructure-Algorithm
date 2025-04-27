package main

import "fmt"

func (ll *linklist) AddRec(n int) {
	newNode := &NodeLL{val: n}
	current := ll.head
	for current.next != nil {
		current = current.next
	}
	ll.head = newNode
	fmt.Println(ll.head.val)
}
func main() {
	ll := &linklist{}
	ll.AddRec(10)
}
