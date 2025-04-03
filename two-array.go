package main

import (
	"fmt"
	"sync"
)

func printInt(a []int, chInt chan bool, chStr chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < len(a); i++ {
		<-chInt
		fmt.Println("The int value: ", a[i])
		chStr <- true
	}
}
func printStr(str []string, chInt, chStr chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < len(str); i++ {
		<-chStr
		fmt.Println("The String value: ", str[i])
		chInt <- true
	}

}

func main() {
	a := []int{1, 2, 3, 4, 5}
	str := []string{"a", "b", "c", "d", "e"}
	chInt := make(chan bool, 1)
	chStr := make(chan bool, 1)
	var wg sync.WaitGroup
	wg.Add(2)
	go printInt(a, chInt, chStr, &wg)
	go printStr(str, chInt, chStr, &wg)
	chInt <- true
	wg.Wait()
	close(chInt)
	close(chStr)
}
