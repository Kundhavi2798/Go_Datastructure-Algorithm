package main

import (
	"fmt"
	"strings"
)

func countVowels(s string) int {
	count := 0
	for _, ch := range s {
		if strings.ContainsRune("aeiou", ch) {
			count++
		}
	}
	return count
}
func reversewordsinstring(s string) {
	str := strings.Fields(s)
	for i := 0; i < len(str)-1; i++ {
		for j := 0; j < len(str)-i-1; j++ {
			str[j], str[j+1] = str[j+1], str[j]
		}
	}
	val := strings.Join(str, " ")
	fmt.Println(val)
}
func findCount(s string) {
	str := strings.Fields(s)
	long := ""
	for _, ch := range str {
		if len(ch) > len(long) {
			long = ch
		}
	}
	fmt.Print(long)
}
func main() {
	fmt.Print(strings.Contains("golang", "go"))
	//find voveles
	str := "education"
	fmt.Println("Vowels Count:", countVowels(str))

	//reverse extra space

	s := "  go   is   fast  "
	res := strings.Join(strings.Fields(s), " ")
	fmt.Print(res)
	reversewordsinstring(res)
	findCount(res)

	//build string repeat
	fmt.Println(strings.Repeat(str, 2))

	var b strings.Builder
	for i := 0; i < 5; i++ {
		b.WriteString("go")
	}
	fmt.Print(b.String())

	data := "Move#Hash#to#Front"
	val := []rune(data)
	a := ""
	count := 0
	for i := 0; i < len(val); i++ {
		if string(s[i]) == "#" {
			count++
			continue
		}
		a += string(s[i])
	}
	hasval := strings.Repeat("#", count)
	fmt.Println(a + hasval)

}


package main

import (
	"fmt"
	"strings"
)

// func main() {
// 	s := "Move#Hash#to#Front"
// 	var builder strings.Builder
// 	count := 0
// 	for _, ch := range s {
// 		if ch == '#' {
// 			count++
// 		} else {
// 			builder.WriteRune(ch)
// 		}
// 	}
// 	builder.WriteString(strings.Repeat("#", count))
// 	fmt.Println(builder.String())
// }
