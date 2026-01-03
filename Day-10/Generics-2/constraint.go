package main

import "fmt"

// Best way: Directly use union of types (no interface needed!)
func Add[T int|float64|int64|float32](a,b T)T{
	return  a+b
}

// if you like named constrained(cleaner for big lists)
type Number interface{
	int|int64|float32|float64|uint|uint64
}

func AddNamed[T Number](a,b T)T{
	return  a+b
}

// if you want + or < operation (like Min/Max)
type OrderedNumber interface{
	~int| ~int64| ~float64| ~float32    //custom types are also allowed
}

func Min[T OrderedNumber](a,b T)T{
	if a<b{
		return  a
	}
	return  b
}

func main() {
	fmt.Println(Add(5, 10))          // 15
    fmt.Println(Add(3.14, 2.86))     // 6
    fmt.Println(AddNamed(100, 200))  // 300

    fmt.Println(Min(10, 20))         // 10
    fmt.Println(Min(5.5, 3.2))    
}