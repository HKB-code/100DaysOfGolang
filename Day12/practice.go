package main

import (
	"fmt"
)

type shoes struct{
	title string
	id string
	price string
	arr []string

}
type product struct{
	title string
	id string
	price float64
}

func main(){
	hobbies:=[3]string{"Reading","Gym","Basketball"}
	fmt.Println(hobbies)
	fmt.Println(hobbies[0])
	newHobbies :=hobbies[0:]
	fmt.Println(newHobbies)
	fmt.Println(hobbies[0:2])
	fmt.Println(hobbies[:2])
	// newHobbies = hobbies[1:3]
	newHobbies = newHobbies[:3]
	fmt.Println(newHobbies)
goals:=[]string{"learn 'go'", "use 'go' for backend"}
fmt.Println(goals)
goals = append(goals, "for cli tools")
fmt.Println(goals)
p:=shoes{
	title: "shoes",
	id: "101",
	price: "2.12",
	arr:[]string{"canvas","spike"},
}
p.arr = append(p.arr, "durable")
fmt.Println(p)
 
ar := []product{
	{
		title: "product-1",
		id: "1",
		price: 1.0,
	},
	{
		title: "product-2",
		id: "2",
		price: 2.0,
	},
}
ar = append(ar, product{
	title: "product-3",
		id: "3",
		price: 3.0,
})
fmt.Println(ar)
}