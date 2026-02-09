package main

import (
	"fmt"
	"time"
)

func main() {
	numofjobs := 20
	numofworkes := 5
	ratelimiter := 5
	jobs := make(chan int, numofjobs)
	worker := make(chan int, numofworkes)

	for i := 0; i < numofjobs; i++ {
		jobs <- i
	}
	close(jobs)

	tokens := make(chan struct{}, ratelimiter)
	for i := 0; i < ratelimiter; i++ {
		tokens <- struct{}{}
	}

	ticker := time.NewTicker(1 * time.Millisecond)
	go func() {
		for range ticker.C {
			select {
			case tokens <- struct{}{}:
			default:
			}
		}
	}()

	assign := func(ids int, job <-chan int, workers chan<- int) {
		for jobs := range job {
			<-tokens
			fmt.Printf("The jobs id %d and processing %d at %v  ", ids, jobs, time.Now())
			fmt.Println()
			time.Sleep(5 * time.Millisecond)
			workers <- jobs
		}

	}

	for i := 1; i <= numofworkes; i++ {
		go assign(i, jobs, worker)
	}

	for i := 1; i < numofjobs; i++ {
		<-worker
	}
	fmt.Print("Processing all the work")
}
