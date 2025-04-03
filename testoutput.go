package main

import "fmt"

func main() {
	slice1 := []int{1, 2}
	slice2 := []int{3, 4}
	slice3 := slice1
	slice1 = slice2
	fmt.Println("1:", slice1)
	fmt.Println("2:", slice2)
	fmt.Println("3:", slice3)

}
