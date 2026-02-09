package main

import (
	"fmt"
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	num := append(nums1, nums2...)
	sort.Ints(num)
	n := len(num)
	res := n / 2
	if n%2 == 1 {
		return float64(num[res])
	} else {
		mid1 := num[res-1]
		mid2 := num[res]
		return float64(mid1+mid2) / 2.0
	}
}
func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	val := findMedianSortedArrays(nums1, nums2)
	fmt.Print(val)
}
