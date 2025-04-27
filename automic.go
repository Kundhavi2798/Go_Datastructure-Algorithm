package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//var jsonData = ``

func main() {
	var counter int32
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt32(&counter, 1)
		}()
	}

	wg.Wait()
	fmt.Println("Final Counter:", atomic.LoadInt32(&counter))
}
