package main

import "fmt"

type LinkList struct {
	val  int
	next *LinkList
}

func buildList(num []int) *LinkList {
	dummy := &LinkList{}
	current := dummy
	for _, num := range num {
		current.next = &LinkList{val: num}
		current = current.next
	}
	return dummy.next
}
func printList(head *LinkList) {
	for head != nil {
		fmt.Print(head.val, "->")
		head = head.next
	}
}
func mergeList(l1 *LinkList, l2 *LinkList) *LinkList {
	dummy := &LinkList{}
	tail := dummy
	for l1 != nil && l2 != nil {
		if l1.val < l2.val {
			tail.next = l1
			l1 = l1.next
		} else {
			tail.next = l2
			l2 = l2.next
		}
		tail = tail.next
	}
	if l1 != nil {
		tail.next = l1
	} else {
		tail.next = l2
	}
	return dummy.next
}
func main() {
	list1 := []int{1, 2, 3, 4}
	list2 := []int{1, 2, 3, 5, 7}
	l1 := buildList(list1)
	l2 := buildList(list2)
	merged := mergeList(l1, l2)
	printList(merged)
}
