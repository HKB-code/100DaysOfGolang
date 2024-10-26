package main

import "fmt"

func main(){
age:=32

var agePointer *int
agePointer = &age
fmt.Println("Age:", age)
fmt.Println("Memmory Address: ", agePointer)
editAge(agePointer)
fmt.Println("Age...",age)
}

func editAge(age*int){
	*age = *age-18
}