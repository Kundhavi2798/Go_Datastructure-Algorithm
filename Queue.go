package main

import "fmt"

type Queue struct {
	value []int
}

func (Q *Queue) enQueue(n int) {
	Q.value = append(Q.value, n)
}

func (Q *Queue) deQueue() int {
	top := Q.value[0]
	Q.value = Q.value[1:]
	return top
}

func main() {
	data := &Queue{}
	data.enQueue(10)
	data.enQueue(20)
	fmt.Println(data.value)
	data.deQueue()
	fmt.Println(data.value)
}
