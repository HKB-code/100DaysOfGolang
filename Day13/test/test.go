package main

import (
	"fmt"
)


type mofify func(int)int

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	doubleValues := transform(&numbers,double)
tripleValues:= transform(&numbers,triple)
	fmt.Println(doubleValues)
	fmt.Println(tripleValues)
moreNumbers:=[]int{5,6,9}

tranformNumbers1:= tranformNumbers(&numbers)
tranformNumbers2:= tranformNumbers(&moreNumbers)

tranformValues1:= transform(&numbers,tranformNumbers1)
tranformValues2:= transform(&moreNumbers,tranformNumbers2)
fmt.Println(tranformValues1)
	fmt.Println(tranformValues2)


}
func transform(number *[]int,m mofify) []int {
	values := []int{}
	for _, val := range *number {
		values = append(values, m(val))
	}
	return values

}

/*
In Go, when you deal with slices, each element in a slice is not a pointer by default, but a value. Therefore, using the dereference operator * on a slice element directly is incorrect because you're trying to dereference something that isn't a pointer.
Value vs Pointer
Values in a Slice: When you access an element from a slice using an index (like numbers[0]), you are accessing the value directly, not a pointer to that value.

Dereferencing Pointers: The dereference operator * is used to access the value that a pointer is pointing to. If numbers[0] is an integer, there's no need to dereference it; you simply use it directly.
*/
// func tranformNumbers(numbers *[]int) mofify{
// 	if *numbers[0]==1{
// 		return double
// 	}else{
// 		return triple
// 	}
// }
func tranformNumbers(numbers *[]int) mofify{
	if (*numbers)[0]==1{
		return double
	}else{
		return triple
	}
}

func double(number int) int {
	return number * 2
}
func triple(number int) int {
	return number * 3
}