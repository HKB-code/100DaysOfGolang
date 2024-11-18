package main

import "fmt"



func main() {
	// different values describing the same thing
	prices := [4]float64{10.99, 9.99, 45.99, 20.00}
	fmt.Println(prices)
	var productNames [4]string
	// array with 4 empty slots
	fmt.Println(productNames)
productNames = [4]string{"A Book"}

//accessing sepicific value
fmt.Println(prices[1])

//setting value
productNames[2] = "A Carpet"

//slice
featuredPrice:= prices[1:3]
fmt.Println(featuredPrice)
highlightedPrice :=featuredPrice[:1]
fmt.Println(highlightedPrice)
/*
When you create a slice,
you don't copy the original array,
so you don't have that copy in memory
that occupies extra memory space.
Instead, you still only have one array in memory
and your slice is just a tiny reference
to a part of that array.
And hence it's a very memory efficient way
of selecting parts of arrays
and if you want to do that
of editing parts of an array.
*/
featuredPrice[0] = 199.99
fmt.Println(prices)

/*
Go also saves some metadata for our slices

that can be useful to look into.

For every slice, we got a length and a capacity,
The length gives us the number

of items in a slice or array.

*/
fmt.Println(len(prices), cap(prices))
fmt.Println(len(highlightedPrice), cap(highlightedPrice))

highlightedPrice = highlightedPrice[:3]
fmt.Println(highlightedPrice)
fmt.Println(len(highlightedPrice), cap(highlightedPrice))


}