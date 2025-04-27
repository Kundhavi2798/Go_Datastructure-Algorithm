package main

import (
	"fmt"
	"time"
)

func printint(s string) {
	fmt.Println(s)
}
func main() {
	str := "Hello"
	go printint(str)
	time.Sleep(time.Second)
	fmt.Println("Kundhavi")

}
