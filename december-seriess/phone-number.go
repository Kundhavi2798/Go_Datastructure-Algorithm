package main

import "fmt"

func Lettercompo(digit string) []string {
	if digit == "" {
		return []string{""}
	}
	phone := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	res := []string{""}
	for i := 0; i < len(digit); i++ {
		var temp []string
		for _, r := range res {
			letter := phone[digit[i]]
			for j := 0; j < len(letter); j++ {
				temp = append(temp, r+string(letter[j]))
			}

		}
		res = temp
	}
	return res
}
func main() {
	digit := "23"
	fmt.Println(Lettercompo(digit))
}
