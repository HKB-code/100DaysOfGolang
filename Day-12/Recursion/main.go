package main

import "fmt"

func main() {
	fact := factorial(5)
	fmt.Println(fact)
	f := fc(5)
	fmt.Println(f)
}

func factorial(number int) int {
	res := 1
	for i := 1; i <= number; i++ {
		res = res * i
	}
	return res
}


func fc(number int )int{
if number==0{
	return  1
}
return  number * fc(number-1)
}