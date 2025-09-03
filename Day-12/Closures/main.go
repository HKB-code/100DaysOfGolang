package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	double := createTranformer(2)
	triple := createTranformer(3)

	doubled := tranformNumbers(&numbers, double)
	tripled := tranformNumbers(&numbers, triple)

	fmt.Println(doubled)
	fmt.Println(tripled)


		transformed := tranformNumbers(&numbers, func(number int) int {
		return number * 2
	})
	fmt.Println(transformed)

}

func tranformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func createTranformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}

/*
🧬 Closures as Families
Imagine a closure as a child function that’s born inside a parent function. The parent has some wealth — variables, parameters, or state — and the child inherits access to that wealth even after the parent has “moved on.”
So yes — closures inherit the money (variables) from their parent scope. Even if the parent function has finished executing, the child (closure) still holds onto those riches.

💡 Go Example: Closure with Inherited Wealth
func multiplier(factor int) func(int) int {
	return func(n int) int {
		return n * factor // 'factor' is inherited from the parent
	}
}

Here:
- factor is the money.
- func(n int) int { return n * factor } is the child.
- Even after multiplier finishes, the child still has access to factor.

🧵 Tweetable Analogy
Closures in Go are like trust fund kids.
Born inside a parent function, they inherit variables like wealth — and keep spending them long after the parent’s gone. 💸
Scope never dies. It just gets captured.



*/