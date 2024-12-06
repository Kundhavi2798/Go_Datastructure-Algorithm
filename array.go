package main

import "fmt"

func sorting(ss []string) {
	length := len(ss)
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-1-i; j++ {
			if ss[j] > ss[j+1] {
				ss[j], ss[j+1] = ss[j+1], ss[j]
			}
		}
	}
}
func bubbleshort(ss []string, low int, high int) {
	if low < high {
		pivotindex := partition(ss, low, high)
		bubbleshort(ss, low, pivotindex-1)
		bubbleshort(ss, pivotindex+1, high)
	}

}
func partition(ss []string, low int, high int) int {
	pivit := ss[high]
	i := low - 1
	for j := low; j < high; j++ {
		if ss[j] > pivit {
			i++
			ss[j], ss[j+1] = ss[j+1], ss[j]
		}
	}
	ss[i+1], ss[high] = ss[high], ss[i+1]
	return i + 1
}
func main() {
	s := []string{"aasd", "bsds", "aa"}
	sorting(s)
	fmt.Println("sorted array", s)
	bubbleshort(s, 0, len(s)-1)
	fmt.Println("sorted bubbleshort", s)

}
