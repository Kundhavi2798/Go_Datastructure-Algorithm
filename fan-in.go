package main

import (
	"fmt"
	"time"
)

func producer(str string, ch chan string) {
	for i := 0; i <= 3; i++ {
		ch <- fmt.Sprintf("the name %s the number of time %d", str, i)
		time.Sleep(time.Millisecond * 200)
	}
	close(ch)
}
func fanin(ch1, ch2 chan string) chan string {
	out := make(chan string)
	go func() {
		for str := range ch1 {
			out <- str
		}
	}()
	go func() {
		for str := range ch2 {
			out <- str
		}
	}()
	return out
}
func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go producer("Producer", ch1)
	go producer("consumer", ch2)
	merged := fanin(ch1, ch2)
	for i := 0; i < 6; i++ {
		fmt.Println(<-merged)
	}
}
