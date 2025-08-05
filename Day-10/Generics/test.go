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
// without generics

// package main

// import "fmt"

// type Instack struct {
// 	items []int
// }
// func (s *Instack) push(item int){
// 	s.items = append(s.items, item)
// }

// func  (s *Instack) pop()int{
// 	item := s.items[len(s.items)-1]
// s.items = s.items[:len(s.items)-1]
// return  item
// }
// func main() {
//     stack := Instack{}
//     stack.push(10)
//     stack.push(20)
// 	fmt.Println(stack)
//     fmt.Println(stack.pop())  // Output: 20
// }

//////////////
// with generics

package main

import "fmt"


type Stack[T any] struct{
	items []T
}

func (s*Stack[T]) Push (item T){
    s.items = append(s.items,item)
}
func (s *Stack[T]) Pop() T {
    item := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return item
}

func main(){
	// Int stack
	intStack := Stack[int]{}
	intStack.Push(10)
	intStack.Push(20)
	  fmt.Println(intStack.Pop())  // Output: 20

    // String stack
	stringStack := Stack[string]{}
    stringStack.Push("Go")
    stringStack.Push("Gophers")
    fmt.Println(stringStack.Pop())  // Output: Gophers

	/*
	🧩 Constraints (in Generics)
Constraints tell Go what types are allowed when someone uses your generic type or function.

	*/
	//🔹 Example: Restricting with comparable
	names := []string{"Harshil", "Aditi", "Rohan", "Neha"}
    index := Find(names,"Rohan")
	if  index!= -1 {
	fmt.Printf("Found 'Rohan' at index %d\n", index)
	} else {
		fmt.Println("Name not found")
	}

}

/*
- T comparable: Only types that support == and != (like numbers, strings, etc.).
- Without constraints, T could be any type—even ones that don’t support ==.

*/
func Find[T comparable](slice []T,target T)int{
	for i,v := range slice{
		if v==target{
			return  i
		}

	}
	return  -1
}