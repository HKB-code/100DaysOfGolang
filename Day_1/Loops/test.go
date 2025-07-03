package main

import "fmt"

func main() {
	// 1. Standard For Loop
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 2. For-Range Loop
	// Iterating over a slice
	num:=[]int{1,2,3,4,5}
	for index,value := range num{
		fmt.Printf("Index: %d, Value: %d\n", index,value)
	}

	// Iterating over a map
	dict:= map[string]int{"a": 1, "b":2,"c":3}
	for key,value := range dict{
		fmt.Printf("Key: %s, Value: %d\n", key,value)


	}

	//3. Infinite loop
	for{
		fmt.Println("Infinite loop")
		break //Avoiding actual infinite loop
	}
	
	//4. Conditional For Loop (While Loop Equivalent)
	n:=1
	for n<10{
		fmt.Println(n)
		n++
	}
}