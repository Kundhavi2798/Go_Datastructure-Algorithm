package main

import (
	"fmt"
	"time"
)

func worker1(ch chan string) {
	time.Sleep(1 * time.Millisecond)
	ch <- "helooo"
}
func worker2(ch chan string) {
	time.Sleep(2 * time.Millisecond)
	ch <- "kundhavi"
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go worker1(ch1)
	go worker2(ch2)
	for i := 0; i < 2; i++ {
		select {
		case name := <-ch1:
			fmt.Println(name)
		case name2 := <-ch2:
			fmt.Println(name2)
		}
	}

}
