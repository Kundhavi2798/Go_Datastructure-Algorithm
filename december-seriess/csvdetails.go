package main

import (
	"time"
	"fmt"
	"os"
)

type Student struct{
	school string
	age int
}
func main(){
	fmt.Println("welcome to coding world")
    var studentParsed []Student

    startTime := time.Now()
	fmt.Println(startTime)
	file ,err := os.Open("/home/kundhavi/student-data.csv")
	if err!=nil{
		fmt.error("csv file open error")
	}
	fmt.Println("csv file open successfully")
	defer file.Close()
     filedetails , err := json.
}