package main

import "fmt"

type shoes struct{
	title string
	id string
	price string
	arr []string
}

type product struct{
	id string
	price int
	title string
}


func main() {

	hobbies := []string{"Gym", "Coding", "Cooking"}
	fmt.Println(hobbies)

	fmt.Println(hobbies[0])

	newHobbies:= hobbies[1:]
	fmt.Println(newHobbies)
fmt.Println(hobbies[:2])

couseGoal := []string{}
couseGoal = append(couseGoal, "Learn Go")
couseGoal = append(couseGoal, "Master Go")
couseGoal = append(couseGoal, "for CLI Tools")
fmt.Println(couseGoal)

p:=shoes{
	title: "Shoes",
	id:"101",
	price: "2.12",
	arr:[]string{"canvas","spike"},
}
p.arr = append(p.arr, "durable")
fmt.Println(p)


arr := []product{
	{title: "Product-1",
id: "1",
price: 1.0,
},
{
title: "Product-2",
id: "2",
price: 2.0,
},


}
arr = append(arr, product{
	title: "Product-3",
id: "3",
price: 3.0,
},)
fmt.Println(arr)
}