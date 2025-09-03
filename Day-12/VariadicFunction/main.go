package main

import "fmt"

func main() {
	number := []int{1, 2, 3, 4, 5}
	sum := sumUp(number...)
	fmt.Println(sum)


	d := s(10,number...)
	fmt.Println(d)
}

func sumUp(numbers ...int) int {
	res := 0
	for _, val := range numbers {
		res += val
	}
	return res
}

func s(firstNumber int , numbers ...int) int {
	res := 0
	for _, val := range numbers {
		res += val
	}
	return res+ firstNumber
}