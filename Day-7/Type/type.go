package main

import "fmt"

/*type in Go:
In Go, the type keyword is used to define custom data types. It gives you the flexibility to create new types tailored to your specific requirements.
The type keyword in Go is used to declare user-defined types (structs, interfaces, and custom primitive types), assign type aliases, and perform type switches at runtime.

Key Uses of type:
1.Structs: Used to create custom composite types.
type Person struct {
    Name   string
    Age    int
    Email  string
}


2.Aliases: Used to give an existing type another name.



3.Custom Types: Used to create your own types that can have specific methods.*/
// type MyString string
// 1.
type Person struct {
	Name  string
	Age   int
	Email string
}

// 2.
type MyString string

// 3.
type Rectangle struct {
	Width  float64
	Height float64
}

// Method to calculate the area of the rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	var ok MyString = "HI"
	fmt.Println("Custom Type: ", ok)
	rect := Rectangle{Width: 10, Height: 5}
	/*When you call rect.Area(), the instance rect is passed as the receiver r to the Area() method. The method then operates on the fields of this instance (rect.Width and rect.Height) to calculate and return the area.*/

	fmt.Println("Area: ", rect.Area())
}
