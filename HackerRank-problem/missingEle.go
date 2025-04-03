package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 5, 6, 8, 10}
	seen := make(map[int]bool)
	res := []int{}
	lastEle := a[len(a)-1]
	fmt.Println(lastEle)
	for _, i := range a {
		seen[i] = true
	}

	fmt.Println("seen", seen)
	for i := 1; i < lastEle; i++ {
		if !seen[i] {
			res = append(res, i)
		}
	}
	fmt.Println(res)
}
