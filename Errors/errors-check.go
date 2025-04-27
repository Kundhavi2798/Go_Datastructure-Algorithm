package main

// import the errors package
import (
	"errors"
	"fmt"
)

type DivisionErros struct {
	error string
}

func (err DivisionErros) Error() string {
	return "Number Cannot Be Divided by Zero"
}

func divide(num1, num2 int) error {
	if num2 == 0 {
		//return fmt.Errorf("the number %d is not a correct dibision", num2)
		return &DivisionErros{}
	}
	return nil
}

func main() {

	message := "Hello"

	myError := errors.New("Number Cannot Be Divided ")

	if message != "ProgramZ" {
		fmt.Println(myError)
	}
	err := divide(4, 0)
	fmt.Println(myError, err)
	data := errors.Unwrap(err)
	fmt.Println(data)
	if err != nil {
		fmt.Println("the error is ", err)
	} else {
		fmt.Println("Correct division")
	}
}
