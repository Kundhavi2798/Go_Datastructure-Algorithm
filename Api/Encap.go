package main

type Employee struct {
	Name string
	Age  int
}

func (Emp Employee) getName() string {
	return Emp.Name
}
func (Emp Employee) getAge() int {
	return Emp.Age
}
func (Emp Employee) setName(name string) {
	Emp.Name = name
}
func (Emp Employee) setAge(age int) {
	Emp.Age = age
}
