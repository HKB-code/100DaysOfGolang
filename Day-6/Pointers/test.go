package main

import "fmt"


func main(){
	age :=32
	var agePointer *int
	agePointer = &age //get the address
	fmt.Println("Age", *agePointer) //get the value (dereference)
result := getAdultYears(agePointer)
fmt.Println(result)

editAge(agePointer)
fmt.Println("mutate" ,age)
}

func getAdultYears(age *int) int{
	return  *age -18
}

func editAge(age *int) {
	*age = *age -18
}