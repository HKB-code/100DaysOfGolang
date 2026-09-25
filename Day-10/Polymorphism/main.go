package main

/*
Polymorphism (meaning "many forms") is the ability of different data types to be treated through a single, unified interface. It allows a single function or component to operate on different concrete types as long as they satisfy a shared contract or constraint.

Unlike traditional object-oriented languages (like Java or C++) that rely on explicit class inheritance (extends / implements), Go implements polymorphism without inheritance, using two primary mechanisms:

1.Subtype Polymorphism via Interfaces (Dynamic / Runtime)

2.Parametric Polymorphism via Generics (Static / Compile-time)
///////////////////////////////////////////////////////////////
1. Subtype Polymorphism via Interfaces (Dynamic)
Go uses implicit interfaces (often called "duck typing with static compile-time checking"). A concrete type implements an interface simply by defining all the methods declared by that interface—no explicit implements keyword is needed.

# How it Works at Runtime:
When you assign a concrete value to an interface variable, Go constructs an internal structure called iface containing:

1.A pointer to the concrete type metadata and method table (itab).

2.A pointer to the underlying data copy.

Method calls on interface variables are resolved dynamically at runtime using this method table (dynamic dispatch).
*/

import (
	"fmt"
	"math"
)

// 1. Define the Interface
type Shape interface {
	Area() float64
}

// 2. Define Concrete Types
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// 3. Polymorphic Function
// Accepts ANY type that satisfies the Shape interface
func PrintArea(s Shape) {
	fmt.Printf("Area: %.2f\n", s.Area())
}

func main() {

	// 1.Subtype:
	shapes := []Shape{
		Circle{Radius: 5},
		Rectangle{Width: 4, Height: 6},
	}

	for _, shape := range shapes {
		PrintArea(shape) // Same function handles both Circle and Rectangle
	}

	// 2.Paramatric
	ints := []int{1, 2, 3, 4}
	floats := []float64{1.5, 2.5, 3.5}

	fmt.Println(Sum(ints))   // Output: 10
	fmt.Println(Sum(floats)) // Output: 7.5
}

////////////////////////////////////////////////////////////////
/*
2. Parametric Polymorphism via Generics (Static)
Introduced in Go 1.18, Generics allow functions, structs, and methods to operate over abstract types using type parameters and constraints.

How it Works at Compile-Time
Unlike interfaces, generics perform static dispatch. The compiler evaluates the concrete types at compile-time and generates optimized code (monomorphization or GCShape dictionary lookup), avoiding runtime interface allocation overhead.
*/

type Numbers interface {
	~int | float64
}

func Sum[T Numbers](slice []T) T {
	var res T
	for _, v := range slice {
		res += v
	}
	return res
}
