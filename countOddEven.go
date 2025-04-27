package main

import "fmt"

func main() {
	arr := []int{22, 32, 42, 52, 62}
	oddCount, evenCount := 0, 0
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 0 {
			evenCount++
		} else {
			oddCount++
		}
	}
	fmt.Printf("oddCount :%d EvenCount : %d\n", oddCount, evenCount)
}
