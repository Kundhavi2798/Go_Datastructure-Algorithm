package main

import "fmt"

func reverseindex(s string, k int) string {
	if k <= 0 {
		return s
	}

	a := []rune(s)
	for i := 0; i <= len(a); i += k {
		fmt.Print(len(a))
		left := i
		right := i + k - 1
		if right >= len(a) {
			right = len(a) - 1
		}
		fmt.Print(left, right)

		for left < right {
			a[left], a[right] = a[right], a[left]
			left++
			right--
		}
	}
	return string(a)
}
func maxSubArray(nums []int) int {
	maxSum := nums[0]
	currSum := 0

	for _, num := range nums {
		if currSum < 0 {
			currSum = 0
		}
		currSum += num
		if currSum > maxSum {
			maxSum = currSum
		}
	}
	return maxSum
}

func main() {
	val := "abcdefghi"
	k := 2
	fmt.Println(reverseindex(val, k))

	// a := []rune(val)

	// for i := 0; i < len(a); i += k {
	// 	left := i
	// 	right := i + k - 1

	// 	if right >= len(a) {
	// 		right = len(a) - 1
	// 	}

	// 	for left < right {
	// 		a[left], a[right] = a[right], a[left]
	// 		left++
	// 		right--
	// 	}
	// }

	// fmt.Println(string(a))
}
