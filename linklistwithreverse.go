// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type NodeLL struct {
	val  int
	next *NodeLL
}

type linklist struct {
	head *NodeLL
}

func (ll *linklist) Add(n int) {
	newNode := &NodeLL{val: n}
	if ll.head == nil {
		ll.head = newNode
		return
	}
	current := ll.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}
func (ll *linklist) Display() {
	if ll.head == nil {
		fmt.Println("Link List is empty")
	}
	current := ll.head
	for current.next != nil {
		fmt.Println(current.val)
		current = current.next
	}

}
func (ll *linklist) ReverseDisplay() {
	stack := []int{}
	if ll.head == nil {
		fmt.Println("Link List is empty")
	}
	current := ll.head
	for current.next != nil {
		stack = append(stack, current.val)
		current = current.next
	}
	for i := len(stack) - 1; i >= 0; i-- {
		fmt.Println(stack[i])
	}

}
func (ll *linklist) remove(n int) {
	fmt.Print(n, ll.head.val)
	if ll.head.val == n {
		ll.head = ll.head.next
	}
	fmt.Println("After remove", ll.head.val)
	current := ll.head
	for current.next != nil && current.next.val != n {
		current = current.next
	}

	current = current.next
}
func main() {
	ll := &linklist{}
	ll.Add(10)
	ll.Add(20)
	ll.Add(30)
	ll.Add(40)
	ll.Display()
	ll.ReverseDisplay()
	ll.remove(10)
	ll.Display()

}
