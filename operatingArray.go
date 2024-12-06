package main

import "fmt"

func searchelm(arr []int, val int) bool {
	for _, num := range arr {
		if num == val {
			return true
		}
	}
	return false
}
func insertelm(arr []int, val1 int, val2 int) bool {
	if val2 < 0 && val2 > len(arr) {
		return false
	}
	arr = append(arr[:val2], append([]int{val1}, arr[val2:]...)...)
	fmt.Println("After insert the array :", arr)
	return true
}
func deleteelm(arr []int, z int) bool {
	for i, val := range arr {
		if val == z {
			arr = append(arr[:i], arr[i+1:]...)
		}
	}
	fmt.Println("After Delete the array :", arr)

	return true
}

func main() {
	arr1 := []int{17, 15, 8, 9, 12}
	x1 := 10
	y1 := 6
	yi1 := 2
	z1 := 5
	fmt.Println("searchelament", searchelm(arr1, x1))
	fmt.Println("Insert element", insertelm(arr1, y1, yi1))
	fmt.Println("Detele element", deleteelm(arr1, z1))
}
