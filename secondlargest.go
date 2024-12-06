package main

import "fmt"

func secondlargest() {
	arr := []int{10, 10, 10}
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] == arr[j+1] {
				arr[j] = -1
			}
		}
	}
	fmt.Println("seconnd largest", arr[1])
}
func movezeroend() {
	arr := []int{0, 1, 9, 8, 4, 0, 0, 2, 7, 0, 6, 0, 9}
	zeromove := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			arr[i], arr[zeromove] = arr[zeromove], arr[i]
			zeromove++
		}
	}
	fmt.Println("lastest array:", arr)
}
func reversearray(arr []int, start int) {
	//arr := []int{1, 4, 3, 2, 6, 5}
	//start := 0
	end := len(arr) - 1
	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start += 1
		end -= 1
	}
	fmt.Println("revesed array", arr)
}
func givenumrotate() {
	arr := []int{1, 4, 3, 2, 6, 5}
	target := 2
	n := target % len(arr)
	fmt.Println("remainder", n, len(arr), arr[n:], arr[:n])
	arr = append(arr[n:], arr[:n]...)
	fmt.Println("reversed array based on target:", arr)

}
func nextPermutation(arr []int) {
	n := len(arr)
	i := n - 2
	fmt.Println("n and i", n, i)
	for i >= 0 && arr[i] >= arr[i+1] {
		fmt.Println("after 1st step", arr[i])
		i--
	}
	fmt.Println("after 1st step", arr)
	fmt.Println("i ", i)

	if i >= 0 { // If such an element is found
		// Step 2: Find the smallest element greater than arr[i]
		j := n - 1
		for arr[j] <= arr[i] {
			j--
		}
		fmt.Println("i and j", i, j)
		// Swap arr[i] and arr[j]
		arr[i], arr[j] = arr[j], arr[i]
	}
	fmt.Println("after swap arr", arr)
	reversearray(arr, i+1)

}

func main() {
	secondlargest()
	movezeroend()
	//reversearray()
	givenumrotate()
	arr := []int{2, 4, 1, 7, 5, 0}
	nextPermutation(arr)
	fmt.Println("Next permutation:", arr)
}
