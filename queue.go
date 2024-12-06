package main

import "fmt"

type Oueue struct {
	val []int
}

func (q *Oueue) queued(n int) {
	q.val = append(q.val, n)
}

func (q *Oueue) dequeued() int {
	if len(q.val) == 0 {
		fmt.Println("queue is empty")
	}
	top := q.val[0]
	q.val = q.val[1:]
	return top
}

func main() {
	que := &Oueue{}
	que.queued(20)
	que.queued(10)
	que.queued(10)
	fmt.Println("queue value", que.val)
	que.dequeued()
	fmt.Println("queue value", que.val)

}
