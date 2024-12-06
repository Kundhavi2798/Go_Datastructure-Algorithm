package main

import "fmt"

func main() {
	s := []int{2, 1, 5, 5, 5, 5, 6, 6, 6, 6, 6}
	count := make(map[int]int)
	for _, freq := range s {
		count[freq]++
	}
	ans := []int{}
	power := len(s) / 3
	for num, val := range count {
		if val >= power {
			ans = append(ans, num)
		}
	}
	if len(ans) == 0 {
		fmt.Println("no occe=orence []")
	} else {
		fmt.Println("ans ", ans)
	}

}
