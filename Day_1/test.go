package main

import "fmt"

func main(){
	fmt.Println("Hello World")
	/*
	The example shows variables of several types, and also that variable declarations may be "factored" into blocks, as with import statements.
	*/
	var(
		a bool = false
		b string = "Hi"
		c int = 1
	)
	fmt.Println(a,b,c)
}