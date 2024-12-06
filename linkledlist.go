package main

import "fmt"

type Node struct {
	value int
	next  *Node
}
type Linkedlist struct {
	head *Node
}

func (ll *Linkedlist) Add(val int) {
	newnode := &Node{value: val}
	if ll.head == nil {
		ll.head = newnode
		return
	}
	current := ll.head
	for current.next != nil {
		current = current.next
	}
	current.next = newnode
}
func (ll *Linkedlist) Display() {
	if ll.head == nil {
		fmt.Println("emtry list")
		return
	}
	current := ll.head
	for current != nil {
		fmt.Printf("%d ->", current.value)
		current = current.next
	}
	fmt.Println("nil")
}
func (ll *Linkedlist) Remove(val int) {
	fmt.Println("val", val)
	if ll.head == nil {
		fmt.Println("list empty not delete")
		return
	}
	fmt.Println("ll.head.value", ll.head.value)
	if ll.head.value == val {
		ll.head = ll.head.next
	}
	current := ll.head
	for current.next != nil && current.next.value != val {
		current = current.next
	}
	if current.next == nil {
		fmt.Println("The value is not in list")
		return
	}
	current.next = current.next.next
}
func main() {
	ll := &Linkedlist{}
	ll.Add(10)
	ll.Add(20)
	ll.Add(30)
	ll.Remove(30)
	ll.Display()
	ll.Add(40)
	ll.Display()
}
