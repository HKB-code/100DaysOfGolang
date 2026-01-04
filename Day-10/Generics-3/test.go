package main

import "fmt"

// Old way (without generics) – Boxing hoti thi

// func PrintAny(values []any){
// 	for _,v:=range values{
// 		fmt.Println(v)
// 	}
// }

// func main() {
// 	nums:= []int{1,2,3,4}
// 	anySlice := make([]any,len(nums))
// 	for i,n:=range nums{
// 		anySlice[i] = n  // boxing
// 	}
// 	fmt.Println(anySlice)
// 	PrintAny(anySlice)
// }

// New way (with generics) – NO BOXING!

func Print [T any](value []T){
	for _,v:= range value{
		fmt.Println(v)
	}
}

func main() {
	nums:= []int{1,2,3,4}
	Print(nums)
	str:= []string{"Hey","Hi","Hello"}
	Print(str)
}

