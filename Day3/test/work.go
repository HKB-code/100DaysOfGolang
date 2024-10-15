/*
What is a Pointer?
A pointer is a variable that stores the memory address of another variable. In Go, pointers are used to reference and manipulate data stored in other variables without copying the data.
or
A pointer is a variable that stores the address of another variable. Essentially, a pointer points to a memory location where the actual data is stored, instead of directly storing the data itself.

Declaring and Using Pointers
To declare a pointer, you use the * operator. To get the address of a variable, you use the & operator.
*/

// Qus. How to use pointers in argument of a function
//Ans. Using pointers as arguments in a function in Go allows you to modify the original value of a variable or efficiently pass large data structures without copying them. Here’s a step-by-step example to illustrate this:

package main

import "fmt"
 func updateValue (num * int){
	*num = 42 //Dereference the pointer and update the value
 }
 func main(){
	var number = 10
	fmt.Println("Before: ", number)

	//Pass the address of the variable to the function
	updateValue(&number)

	fmt.Println("After: ", number)

///////
var ptr* int = &number
//some memory address
fmt.Println(ptr)
//value
fmt.Println(*ptr)


 }
 /*
 Understanding & and * Operators in Pointers:
 1.& Operator:

a.The & operator is used to pass the address of a variable.

b.When you write &number, it provides the memory location where number is stored.

c.Think of it like giving the address of a house.

2.* Operator:

a.The * operator is used to dereference a pointer.

b.This means it accesses the value stored at the memory address the pointer points to.

c.Without *, you would be working with the address itself, not the value.

d.It's like using a key (*) to open a door at the address and access what's inside

notes:Benefits of Using Pointers in Function Arguments
Efficiency: Avoids copying large data structures.
Mutability: Allows functions to modify the original variable.
Clarity: Makes it clear that the function can change the input variable.
 */