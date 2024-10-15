package main

import "fmt"
func Add(a, b int) int {
	return a + b
}

func main(){
	 var a int = 1000  // Explicitly declared as int
    var b float64 = 1.0  // Explicitly declared as float64
    var c float64 = 5.5  // Explicitly declared as float64

    result := float64(a) + b + c
    fmt.Println(result)  // Output: 1006.5
}