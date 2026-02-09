package main

import "fmt"

type LinkList struct {
	Val  int
	Next *LinkList
}

func sliceToLinkedList(a []int) *LinkList {
	dummy := &LinkList{}
	carr := dummy
	for _, v := range a {
		carr.Next = &LinkList{Val: v}
		carr = carr.Next
	}
	return dummy.Next
}
func addTwoNumbers(l1 *LinkList, l2 *LinkList) *LinkList {
	carry := 0
	dummy := &LinkList{}
	carr := dummy
	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		carr.Next = &LinkList{Val: sum % 10}
		carry = sum / 10
		carr = carr.Next

	}
	return dummy.Next
}
func linklisttoarray(head *LinkList) []int {
	res := []int{}
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}

func main() {
	num := []int{9, 9, 9, 9, 9, 9, 9}
	num2 := []int{9, 9, 9, 9}

	l1 := sliceToLinkedList(num)
	l2 := sliceToLinkedList(num2)
	arr := addTwoNumbers(l1, l2)
	data := linklisttoarray(arr)
	fmt.Print(data)
}
