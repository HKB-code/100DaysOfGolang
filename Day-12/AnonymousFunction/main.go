package main

import "fmt"

type Person struct {
	Name  string
	Age   int
	Email string
}

func (p *Person) Greet(retiermentYear func(int) int) string {
	age := retiermentYear(p.Age)
	return fmt.Sprintf("Hello %v, you have %v years left for Retierment",p.Name,age)
}

func main() {

p:=Person{
		Name: "Jhon",
		Age: 42,
		Email:"JhonDev@gmail.com",
	}
	fmt.Println(p.Greet(func(age int)int{
		return 65 - age
	}))

	numbers := []int{1, 2, 3, 4}

	// Anonymous Function : that anonymous function is a value, not a type.
// - This is a function literal — a value that represents a function. It’s defined inline and passed directly as an argument.

	transformed := transformNumbers(&numbers, func(number int) int {
		return number * 2
	})
	fmt.Println(transformed)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform((val)))
	}
	return dNumbers
}