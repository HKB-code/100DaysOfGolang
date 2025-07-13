package main

import "fmt"

type animal struct {
	name string
}

func (a animal) speak() string {
	return "My name is " + a.name
}

type dog struct {
	animal
	breed string
}

func main() {
	d := dog{
		animal: animal{name: "Buddy"},
		breed:  "Labrador",
	}
	fmt.Println(d.speak())
}