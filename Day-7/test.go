// custom types and adding method
package main

import (
	"fmt"
	"reflect"
)


type str string

func (text str) log(){
	fmt.Println(text)
}

func main(){
	var ok str = "Hello"
	ok.log()

	var z string = "sjkj"
	// Type conversion needed
	str(z).log()
	/*
	-Values of type str are not assignable to string without conversion
	var s str = "hello"
	var x string = s      // ERROR
    var x string = string(s)  // OK

	But they behave exactly like strings internally (same memory layout, same operations).
	*/
t:=reflect.TypeOf(ok)
fmt.Println(t)
fmt.Println(t.Kind())
/*
Yes — the concrete type (string) is what Go uses when the value is stored, passed, or operated on at runtime.
The declared type (str) is only for compile-time checks, type safety, and attaching methods.


The concrete type is the real, actual type that is used when a value is stored or manipulated at runtime.
NOTES:
Why This Matters
At runtime, Go treats a value of type str exactly like a string:
Same memory layout
Same performance
Same internal representation

That's why the concrete type is string.
But at compile time, Go sees str as a different type:
You cannot pass str to a function expecting string without conversion
You can attach methods only to str, not to string


*/
}

/*
Summary
type str string → New custom type (distinct type with underlying type string)
You can add methods to it
Not compatible with string without explicit conversion
Use this when you want type safety or specific behavior

Compile time: Strict distinction (str is separate) → safety + methods.
Runtime: Concrete type (string) → efficient storage and execution.

Why This Design?
Compile time: Gives type safety and ability to add methods/behavior specific to str.
Runtime: No penalty — the custom type "disappears," and it's treated as the concrete type.


*/