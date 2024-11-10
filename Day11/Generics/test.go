//without generic:
/*
package main

import "fmt"

func printInt(value int) {
    fmt.Println(value)
}

func printString(value string) {
    fmt.Println(value)
}

func main() {
    printInt(42)
    printString("Hello, Go!")
}
*/
// with generic:
/*
Generics allow you to write functions and data structures that can operate on any data type. This helps in writing more flexible and reusable code.
*/
/*
package main

import "fmt"

// Generic Print function
func Print[T any](value T) {
    fmt.Println(value)
}

func main() {
    Print(42)            // Works with int
    Print("Hello, Go!")  // Works with string
    Print(3.14)          // Works with float
}
*/
/*
Explanation:

[T any]: Declares a type parameter T that can be any type.

Print Function: Can now print values of any type.

Benefits:
Reusability: Write code once and use it for any type.

Type Safety: The compiler ensures type correctness, preventing errors.

Maintainability: Less code duplication means easier maintenance.
*/

// ////////////////////////////////////////
//withot generics
/*
package main

import "fmt"

type IntStack struct {
    items []int
}

func (s *IntStack) Push(item int) {
    s.items = append(s.items, item)
}

func (s *IntStack) Pop() int {
    item := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return item
}

func main() {
    stack := IntStack{}
    stack.Push(10)
    stack.Push(20)
    fmt.Println(stack.Pop())  // Output: 20
}
*/
//with generic

package main

import "fmt"

// Generic Stack
type Stack[T any] struct {
    items []T
}

func (s *Stack[T]) Push(item T) {
    s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() T {
    item := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return item
}

func main() {
    // Int Stack
    intStack := Stack[int]{}
    intStack.Push(10)
    intStack.Push(20)
    fmt.Println(intStack.Pop())  // Output: 20

    // String Stack
    stringStack := Stack[string]{}
    stringStack.Push("Go")
    stringStack.Push("Gophers")
    fmt.Println(stringStack.Pop())  // Output: Gophers
}
