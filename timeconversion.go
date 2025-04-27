package main

import (
	"fmt"
	"strconv"
	"strings"
)

func convertTo24HourFormat(time12 string) string {
	period := strings.ToUpper(time12[len(time12)-2:])
	timePart := time12[:len(time12)-2]

	parts := strings.Split(timePart, ":")
	hour, _ := strconv.Atoi(parts[0])
	minute := parts[1]
	second := parts[2]
	fmt.Println(second)

	if period == "AM" {
		if hour == 12 {
			hour = 0
		}
	} else if period == "PM" {
		if hour != 12 {
			hour += 12
		}
	}

	return fmt.Sprintf("%02d:%s:%s", hour, minute, second)
}

func main() {
	time12 := "11:20:50PM"
	time24 := convertTo24HourFormat(time12)
	fmt.Println("24-hour format:", time24)
}
