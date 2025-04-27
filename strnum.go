package main

import (
	"fmt"
	"sync"
)

func number(n []int, wg *sync.WaitGroup, chInt, chStr chan bool) {
	wg.Done()
	for i := 0; i < len(n); i++ {
		<-chInt
		fmt.Println(i)
		chStr <- true
	}
}
func stringprint(str []string, wg *sync.WaitGroup, chInt, chStr chan bool) {
	wg.Done()
	for i := 0; i < len(str); i++ {
		<-chStr
		fmt.Println(str[i])
		chInt <- true
	}
}

func main() {
	num := []int{1, 2, 3}
	str := []string{"a", "b", "c"}
	chInt := make(chan bool)
	chStr := make(chan bool)
	var wg sync.WaitGroup
	wg.Add(2)
	go number(num, &wg, chInt, chStr)
	go stringprint(str, &wg, chInt, chStr)
	chInt <- true

	wg.Wait()
	close(chInt)
	close(chStr)

}
