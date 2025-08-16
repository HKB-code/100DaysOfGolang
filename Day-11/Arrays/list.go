package main

import "fmt"


func main(){
	// Array 
	//different values describing the same thing
	prices := [4]float64{10.99,9.99,45.99,20.00}
	fmt.Println(prices)
    var productNames [4]string
    // array with 4 empty slots
	fmt.Println(productNames)
	productNames = [4]string{"A Book"}
	fmt.Println(productNames)
    
//accessing sepicific value
fmt.Println(prices[1])

//setting value
productNames[2] = "A Carpet"
	
//slice
featuredPrice:= prices[1:3]
fmt.Println(featuredPrice)
highlightedPrice :=featuredPrice[:1]
fmt.Println(highlightedPrice)
fmt.Println(prices)
featuredPrice[0] = 199.99
fmt.Println(prices)
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
of editing parts of an array.,
//////////////////////////////
Slices are like a reference,
like a window into an array, a bit like a pointer,
though it's a different concept.
When we create an array like prices,
that array is stored in memory.
When we then treat a slice based on that array,
we get a window into that array, so to say.
And therefore, if you would modify an element in a slice,
we would also modify the same element
in the original array.
*/

/*
Go also saves some metadata for our slices

that can be useful to look into.

For every slice, we got a length and a capacity,
The length gives us the number

of items in a slice or array.

*/
fmt.Println(len(prices), cap(prices))

/*
In Go, a slice doesn't create a copy of the elements; it points to a section of an underlying array. This allows slices to be resized within the bounds of the original array's capacity.
*/
fmt.Println(len(highlightedPrice), cap(highlightedPrice))
// When extending highlightedPrice, it reuses the capacity of the original prices array. Even though highlightedPrice initially had one element, it can access and include elements within the capacity of the underlying array
highlightedPrice = highlightedPrice[:3]
fmt.Println(highlightedPrice)
fmt.Println(len(highlightedPrice), cap(highlightedPrice))
copy := prices[1:3]
fmt.Println(len(copy), cap(copy))

// Building dynamic lists with slices
 
/*
One downside of using that array was that we had to define
ahead of time how many values we'll have in there.
Now, sometimes in the applications you're going to write,
you know exactly how many values you'll have
in a list of values.
For example, if you handle some user input in a form
that has exactly three form fields,
you might have an easy time storing those three values
in an array because you know ahead of time
that you will get three values.
But let's say you are loading a list
of prices from a database,
and the data in that database changes frequently.
Then when you write your code, you don't know ahead
of time whether you'll need four or five or 100 prices.
So having to set the length of an array in advance
in your code can be a restriction that's quite annoying.
In other programming languages like JavaScript,
you don't have that restriction.
There, you can create an array,
and it's automatically dynamic, so you can always add
and remove elements just as you want to.
For Go arrays, that's not the case.
Prices is an array with four elements,
and you can't just add a fifth element,
even if you wanted to do that.
Now, since we have that restriction,
Go thankfully also has a solution for us,
and that solution again is that slices feature.
*/
amount := []float64{10.00,9.00}
fmt.Println(amount[1])
fmt.Println(amount[0:1])
amount[1] = 10.999
// So we still can't access indexes that don't exist,
// amount[2] = 11.00
/*
but I told you that this would be dynamic, right?
But not like this.
Instead, there now is a built-in function in Go,
the append function
*/
 amount = append(amount, 5.99)

fmt.Println(amount)
/*
You can still use fixed length arrays if you know for sure
that you will never have more than X amounts of elements
in a list, but if there is at least a slight chance
that you might want to add new elements
as your program executes, you should definitely go
for this slices approach right from the start instead,
because that allows you to add elements to slices,
and to let your arrays grow just as they need to.
*/
// for removing:We can of course just create a new slice that removes the element we wanna remove.
amount = amount[1:]
fmt.Println(amount)
//unpack slice
originalPrice := []float64{1.22,2.22,3.4,10.99}
discountPrices :=[]float64{1.1,2.1,3.1,5.1}
originalPrice =append(originalPrice, discountPrices...)
// ...: The variadic unpacking operator—it expands the slice into individual elements.

fmt.Println(originalPrice)
}

// Summary:
/*
What is a Slice?
A slice is a flexible and dynamic data structure in Go that allows you to work with sequences of elements. Unlike arrays, slices can grow and shrink as needed

Creating a Slice
You can create a slice in Go without specifying a fixed size.

package main

import "fmt"

func main() {
    // Creating a slice with initial values
    amount := []float64{10.00, 9.00}
    fmt.Println("Initial slice:", amount)
}
Adding Elements to a Slice
To add more elements to a slice, you use the append function. This function dynamically resizes the slice if needed, by creating a new underlying array and copying the existing elements into it.

package main

import "fmt"

func main() {
    // Creating a slice with initial values
    amount := []float64{10.00, 9.00}
    fmt.Println("Initial slice:", amount)

    // Adding more elements to the slice
    amount = append(amount, 12.50, 7.75)
    fmt.Println("Updated slice:", amount)

    // Adding another single element to the slice
    amount = append(amount, 15.00)
    fmt.Println("Final slice:", amount)
}
How it Works Behind the Scenes
When you create a slice, Go manages an underlying array for you. If you keep appending elements and exceed the capacity of the initial array, Go automatically creates a new, larger array, copies the existing elements over, and then appends the new elements. This process is efficient and allows for dynamic resizing.

Capacity and Length
Slices in Go have a length (the number of elements in the slice) and a capacity (the number of elements the slice can hold before needing to allocate more memory).

*/