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

	// Note:
	/*
	Example ❌
	if str := "Hello"; str {
    fmt.Println(str)
}

- str := "Hello" is fine.
- But the condition part is just str.
In Go, unlike some languages (like C or Python), a string cannot be directly used as a boolean.
The compiler will throw an error:
- Go enforces strict typing: the condition must be explicitly a bool.

Why Go enforces this
- Go avoids implicit truthiness (like empty string = false, non-empty = true).
- This design choice makes code more predictable and less error-prone.
- If you want to check whether a string is non-empty, you must write:

if str := "Hello"; str != "" {
    fmt.Println(str)
}

👉 So the rule is:
- You can declare a variable before the semicolon in an if.
- But the condition after the semicolon must be a boolean expression, not just a value of another type.

	*/
}