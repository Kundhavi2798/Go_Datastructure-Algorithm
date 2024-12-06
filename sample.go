package main

import (
	"fmt"
)

func print(n int) {

	fmt.Println("number printing", n)

}

func main() {
	fmt.Println("go routine")
	print(5)
	channel := make(chan int, 2)
	channel <- 1
	channel <- 2
	fmt.Println("channel1", <-channel)
	fmt.Println("channel2", <-channel)
	//time.Sleep()
}
