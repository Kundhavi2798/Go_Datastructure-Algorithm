package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func maxSubArray(a []int) int {
	maxsum := a[0]
	cursum := a[0]

	for _, i := range a[1:] {
		if cursum < 0 {
			cursum = i
		} else {
			cursum += i
		}

		if cursum > maxsum {
			maxsum = cursum
		}
	}
	return maxsum
}
func main() {

	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println("Max Subarray Sum:", maxSubArray(arr))
	sampleFile := "text.txt"
	file, err := os.Create(sampleFile)
	if err != nil {
		log.Fatal("error", err)
	}
	fileData, err := file.WriteString("hello kundhavi")
	if err != nil {
		log.Fatal("error", err)
	}
	fmt.Println(fileData)
	readData, err := ioutil.ReadFile(sampleFile)
	if err != nil {
		log.Fatal("error", err)
	}
	fmt.Println(string(readData))
}
