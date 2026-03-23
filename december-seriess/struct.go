package main

import "fmt"

type student struct {
	Name string
	Age  int
}

func main() {
	rollnumber := make(map[int]student)
	rollnumber[101] = student{Name: "kundhavi", Age: 28}
	rollnumber[102] = student{Name: "mayavathi", Age: 29}
	roll := 101
	if s, ok := rollnumber[roll]; ok {
		fmt.Println("Name:", s.Name, "Marks:", s.Age)
	} else {
		fmt.Println("Student not found")
	}
}
