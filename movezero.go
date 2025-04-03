package main

import "fmt"

func main() {
	arr := []int{1, 2, 0, 4, 3, 0, 5, 0}
	n := len(arr)
	zero := 0
	for i := 0; i < n; i++ {
		if arr[i] != 0 {
			arr[i], arr[zero] = arr[zero], arr[i]
			zero++
		}
	}
	fmt.Println(arr)
}
