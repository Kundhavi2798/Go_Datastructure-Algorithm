package main

import "fmt"

type Branch struct {
	BName string
}
type mens struct {
	Branch
	Name string
	Age  int
}

type womens struct {
	Branch
	Name string
	Age  int
}

func (M mens) printMens() {
	fmt.Println("The details", M.BName, M.Name, M.Age)
}
func (W womens) printWomens() {
	fmt.Println("The details", W.BName, W.Name, W.Age)
}

func main() {
	men := mens{Branch{"Computer"}, "bhash", 20}
	Womens := womens{Branch{"Computer"}, "kundhavi", 20}
	men2 := mens{Branch{"Mech"}, "Sakthi", 27}
	men.printMens()
	Womens.printWomens()
	men2.printMens()
}
