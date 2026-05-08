package main

import (
	"fmt"
	"sync"
)

func worker(job chan int, res chan int, i int, wg *sync.WaitGroup) {
	defer wg.Done()
	for jobs := range job {
		fmt.Printf(" the processing job %d and id %d  ", i, jobs)
		result := jobs * 2
		res <- result
	}
}

func main() {
	totaljob := 1000
	totalworker := 100
	job := make(chan int, totaljob)
	res := make(chan int, totaljob)
	var wg sync.WaitGroup
	for i := 0; i < totalworker; i++ {
		wg.Add(1)
		go worker(job, res, i, &wg)
	}
	for i := 0; i < totaljob; i++ {
		job <- i
	}
	close(job)

	wg.Wait()
	close(res)

	count := 0
	for result := range res {
		fmt.Println("the result", result)
		count++
	}
	fmt.Println("The processed count : ", count)
}
