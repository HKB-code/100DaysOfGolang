package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) float64 {
	if x == 0 {
		return 0
	}
	z := x / 2
	const epsilon = 1e-12
	fmt.Printf("Initial Guess: %.15f\n", z)

	iteration:=0
	for{
		iteration++
		zNew := z-(z*z-x)/(2*z)

		difference := math.Abs(zNew-z)
		fmt.Printf("Iteration %2d:  %.15f   (change: %.2e)\n", 
            iteration, zNew, difference)

			// stop
			if difference<epsilon{
				 fmt.Printf("→ Stopped after %d iterations (change < %g)\n", 
                iteration, epsilon)
            break
			}
			z=zNew
			if iteration > 100 {
            fmt.Println("→ Safety limit reached (100 iterations)")
            break
        }
	}
	return  z
	
}

func main() {
   fmt.Println("Computing square root of 25:")
    result := sqrt(25)
    fmt.Printf("\nFinal result: %.15f\n", result)
    fmt.Println("Expected:     5.000000000000000")
}