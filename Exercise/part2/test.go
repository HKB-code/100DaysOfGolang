/*
Exercise: Loops and Functions
As a way to play with functions and loops, let's implement a square root function: given a number x, we want to find the number z for which z² is most nearly x.

Computers typically compute the square root of x using a loop. Starting with some guess z, we can adjust z based on how close z² is to x, producing a better guess:

z -= (z*z - x) / (2*z)
*/
package main

import "fmt"

func sqrt(x float64)float64{
	z := x/2 // initial guess
	for i:=0;i<10;i++{
		z-=(z*z -x) /(2*z)
		fmt.Println(z)
	}
	return  z
}
func main() {
	fmt.Println(sqrt(25))
}

/*
Note:
1. z-=(z*z-x)/2*z
- z*z - x → the error term (how far z² is from the target x).
- (z*z - x) / (2*z) → the Newton correction step for the function f(z)=z^2-x.
- z -= ... → update z by subtracting the correction, nudging it closer to the true square root.


2. initial guess : z/2
- We need an initial guess z to start Newton–Raphson iteration:
- The choice of z doesn’t have to be perfect; Newton’s method will converge as long as the guess is positive and not zero.
- x / 2 is a simple heuristic:
- For large x, x/2 is closer to √x than starting at 1.
- For small x, it’s still a positive number that keeps the iteration stable.

3.🎯 Why x/2 is chosen
- Balance between simplicity and convergence speed:
- If x = 100, √x = 10. Starting at x/2 = 50 is not exact, but Newton’s method quickly pulls it down toward 10.
- If x = 4, √x = 2. Starting at x/2 = 2 is already perfect.
- Avoid division by zero:
- Newton’s formula divides by 2z. If you started with z = 0, you’d crash. x/2 guarantees a safe, nonzero start.
- General-purpose guess:
- Works for any positive x without special casing.
- Keeps the algorithm simple and deterministic.

4.🛠 Backend analogy for you
Think of z := x/2 like setting a default shard size in a distributed DB:
- It’s not optimal for every dataset, but it’s safe, predictable, and gets you started.
- The system (Newton’s iteration) then rebalances automatically until it converges to the right distribution (the true √x).

*/