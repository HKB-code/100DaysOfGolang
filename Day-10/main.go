// package main

// import "fmt"

// func add(a, b interface{}) interface{} {

// 	aInt, isaInt := a.(int)
// 	bInt, isbInt := b.(int)
// 	if isaInt && isbInt {
// 		return aInt + bInt
// 	}
// 	aFloat, isaFloat := a.(float64)
// 	bFloat, isbFlaot := b.(float64)
// 	if isaFloat && isbFlaot {
// 		return aFloat + bFloat
// 	}
// 	aStr, isaStr := a.(string)
// 	bStr, isbStr := b.(string)
// 	if isaStr && isbStr {
// 		return aStr + bStr
// 	}

// 	return nil
// }

// func main() {
// 	result := add(1, 2)
// 	fmt.Println(result)
// }

// we can make this code more generic....

package main

import "fmt"

func add [T int|float64|string](a,b T)T{
	return  a+b
}
func main(){
	result:=add(1,2)
	fmt.Println(result)
}
/*
Now it's worth noting that you might not need this feature

here too often

because if you are writing your own programs,

you don't often need code that is that generic.

But especially if you are working on a library

that should be shared with other people,

where you don't necessarily know in advance

with which kinds of values they might be using

your functions, especially in cases like this,

this generics feature here can be really useful.*/