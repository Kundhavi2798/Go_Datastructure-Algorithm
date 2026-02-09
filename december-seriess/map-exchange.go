package main

import "fmt"

func main() {
	mapping := map[string]int{
		"apple":  5,
		"banana": 3,
		"cherry": 7,
	}
	res := make(map[int]string)

	for i,j:=range mapping{
		res[j]= i
	}
	fmt.Println(res)
}