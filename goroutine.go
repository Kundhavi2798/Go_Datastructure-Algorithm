package main

import "fmt"

func main() {
	input := []int{1, 2, 2, 2, 3, 3, 4, 4, 4, 4, 4}
	count := make(map[int]int)
	for _, freq := range input {
		count[freq]++
	}
	for _, freq := range count {
		freq++
	}
	for num, freq := range count {
		fmt.Println("%d number occured %d", num, freq)
	}

	//count := make(map[int]int)
	//for _, char := range input {
	//	count[char]++
	//}
	//freq := []int{}
	//for _, frequencies := range count {
	//	freq = append(freq, frequencies)
	//}
	//fmt.Println("freq", freq)
	//for i := 0; i < len(freq)-1; i++ {
	//	for j := 0; j < len(freq)-1-i; j++ {
	//		if freq[j] < freq[j+1] {
	//			freq[j], freq[j+1] = freq[j+1], freq[j]
	//		}
	//	}
	//}
	//fmt.Println("freq tes", freq)
	//if len(freq) < 2 {
	//	fmt.Println("occured no freq")
	//}
	//secoundfreq := freq[1]
	//for num, freq := range count {
	//	if freq == secoundfreq {
	//		fmt.Println("%d occured %d num", num, secoundfreq)
	//	}
	//}
}
