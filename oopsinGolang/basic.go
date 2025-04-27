package main

import "fmt"

// class, object
type Employee struct {
	Name string
	Age  int
}

func (Emp Employee) printDetails() {
	fmt.Printf("The Employee name %s and Age %d\n", Emp.Name, Emp.Age)
}
func main() {
	Emp1 := Employee{Name: "kundhai", Age: 30}
	Emp2 := Employee{Name: "Maya", Age: 20}
	Emp1.printDetails()
	Emp2.printDetails()
}
