// package main

// import "fmt"

//	func main() {
//		num := -123
//		c := fmt.Sprintf("%d", num)
//		res := ""
//		if num <= 0 {
//			fmt.Print("regative")
//			res = "-"
//			c = c[1:]
//		}
//		for i := len(c) - 1; i >= 0; i-- {
//			res += string(c[i])
//		}
//		fmt.Print(res)
//	}
package main

import "fmt"

func main() {
	num := -112345
	rev := 0

	for num != 0 {
		digit := num % 10
		rev = rev*10 + digit
		num = num / 10
	}

	fmt.Println("Reversed:", rev)
}
