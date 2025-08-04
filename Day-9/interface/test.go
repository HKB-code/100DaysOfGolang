/*
// 🧩 What is an Interface in Go?
// An interface defines a set of method signatures. Any type that implements those methods implicitly satisfies the interface — no implements keyword needed.
*/
package main

import (
	"fmt"
	"math"
)
type Shape interface {
    Area() float64
    Perimeter() float64
	/*
	If a Circle or Rectangle type defines both Area() and Perimeter(), it satisfies the Shape interface.

	*/
}
/*
🔍 Key Concepts
- Implicit Implementation: No need to declare that a type implements an interface — it just needs to have the required methods.
- Polymorphism: Interfaces allow different types to be treated uniformly.
- Dynamic Dispatch: Method calls on interface variables are resolved at runtime.

*/
type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
    return 2 * math.Pi * c.Radius
}

func main(){
	var s Shape = Circle{Radius: 5}
fmt.Println(s.Area())      // Works because Circle satisfies Shape


// 🧠 Empty Interface (interface{})
// The empty interface can hold any value — it's Go’s version of any.
var x interface{}
x = 42
x = "hello"
fmt.Println(x)


/*
🧪 Type Assertion
Extract the concrete type from an interface:

*/
var i interface{} = "hello"
z := i.(string) 
fmt.Println(z)    // z is now a string
z, ok := i.(string)
if(ok){
	fmt.Println(z)
}

/*
🔄 Type Switch
Handle multiple types dynamically:

*/
switch v := i.(type) {
case int:
    fmt.Println("int:", v)
case string:
    fmt.Println("string:", v)
default:
    fmt.Println("unknown type")
}
}

/*
🧱 Interface Composition
Interfaces can be composed:

*/
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

type ReadWriter interface {
    Reader
    Writer
}
/*
💡 Real-World Use Cases
- io.Reader, io.Writer: foundational for file and network operations
- http.Handler: core to building web servers
- Custom interfaces for mocking in tests

*/