package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	data := "2:50:00 AM"
	lastdata := data[len(data)-2:]
	datedata := data[:len(data)-2]
	t := strings.Split(datedata, ":")
    date,_ := strconv.Atoi(t[0])
    min := t[1]
    sec := t[2]
    if lastdata =="PM" && date !=12{
        date+=12
    } else if lastdata == "AM" && date == 12{
        date =0
    }
    hours := fmt.Sprintf("%02d", date)
    fmt.Println(hours + ":"+ min +":"+ sec )
}
