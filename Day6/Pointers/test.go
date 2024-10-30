package main

import "fmt"


type Person struct{
	name string
	age int
}

func main(){
age:=32

var agePointer *int
agePointer = &age
fmt.Println("Age:", age)
fmt.Println("Memmory Address: ", agePointer)
editAge(agePointer)
fmt.Println("Age...",age)



//1.Direct Access: You can directly access fields without pointers.
john:=Person{
	name: "John",
	age: 30,
} 
fmt.Println(john.name,john.age)//Direct accesss

//2.Pointer to a Struct: You can create a pointer to a struct and access, its fields using the pointer.
jhonPointer:=&john
fmt.Println(jhonPointer.name)//Acess via Pointer

//3.Modifying Struct via Pointer: Pointers allow you to modify the original struct.
jhonPointer.age = 31
fmt.Println(john.age)//output:31

/*
Structs: Used for grouping related data.

Pointers: Can reference structs to access or modify data.

Interplay: Pointers offer a way to work with structs more dynamically
*/


}

func editAge(age*int){
	*age = *age-18
}