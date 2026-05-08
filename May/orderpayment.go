package main

import "fmt"
type Order struct{
	id int
	val int
}
type Payments struct{
	id int
	val int
}
func main(){
    order := []Order{
		{id:1, val: 10},
		{id :2, val:30},
		{id : 3, val :40},
	}
	payment := []Payments{
		{id:1, val:10},
		{id:2,val:10},
		{id:2, val:20},
		{id:3,val:30},
		{id:4,val:40},
	}
	ordermap := make(map[int]int)
	for _,data := range order{
		ordermap[data.id]= data.val
	}
	fmt.Println(ordermap)
	paymentmap :=make(map[int]int)
	for _,data := range payment{
		paymentmap[data.id] += data.val
	}
	fmt.Println(paymentmap)
	for id,orderval := range ordermap{
		paymentval, exsit := paymentmap[id]
		if !exsit{
          fmt.Println("Missing Order")
		}else if paymentval==orderval{
			fmt.Println(id,"Matched")
		}else{
			fmt.Println(id,"Not Matched")
		}
	}
	for id, _ := range paymentmap{
	    if _,exsit:= ordermap[id]; !exsit{
			fmt.Println(id,"Misisng Payments")
		}
	}
}