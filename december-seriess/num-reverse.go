package main

import (
	"fmt"
	"strconv"
)

func reverval(input float64) float64 {
	sign := 1.0
	if input < 0 {
		sign = -1
		input = -input
	}
	fmt.Println(input)
	val := strconv.FormatFloat(input, 'f', -1, 64)
	strval := []rune(val)
	for i := 0; i < len(strval)-1; i++ {
		for j := 0; j < len(strval)-i-1; j++ {
			strval[j], strval[j+1] = strval[j+1], strval[j]
		}
	}
	againFloat, err := strconv.ParseFloat(string(strval), 64)
	if err != nil {
		return 0
	}
	return againFloat * sign

}
func main() {
	var num float64 = -12345.67
	val := reverval(num)
	fmt.Println(val)
}