package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func worker(id int, ch1 <-chan int, ch2 chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for jobs := range ch1 {
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		fmt.Println("The process id", jobs, id)
		ch2 <- jobs * 2
	}
}

func main() {
	ch1 := make(chan int, 5)
	ch2 := make(chan int, 5)
	var wg sync.WaitGroup

	for i := 0; i <= 3; i++ {
		wg.Add(1)
		go worker(i, ch1, ch2, &wg)
	}
	for i := 0; i <= 5; i++ {
		ch1 <- i
	}
	close(ch1)

	go func() {
		wg.Wait()
		close(ch2)
	}()
	for res := range ch2 {
		fmt.Println("the result :", res)
	}
}
