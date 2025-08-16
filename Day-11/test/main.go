package main

import "fmt"

// using type aliases
type floatMap map[string]float64

func (m floatMap) output(){
	fmt.Println(m)
}

func main() {
	// this will create an error
	/*
	🔥 Why It Fails
- You're trying to assign to index 0, but the slice has no elements.
- []string{} creates a slice with length 0—you can't assign to an index that doesn't exist.

	*/
	// userName:=[]string{}
	// userName[0] = "Hi"
	// fmt.Println(userName)

	/*
			The make() function in Go is a built-in function used to create and initialize slices, maps, and channels. It allocates and initializes an object of the specified type, ready for use, and returns the object.

			Purpose of make():
		Slices: make() is used to create slices with a specific length and capacity.

		Maps: make() is used to create maps with an initial capacity.

		Channels: make() is used to create channels with a specified buffer size.



		Why make()?
		Slices, Maps, and Channels: These are reference types that require initialization before use. make() sets up the internal data structures needed for these types to function correctly.

		Memory Allocation: Ensures that the memory is allocated appropriately and the object is in a usable state.

		Performance: Allows you to specify initial capacities, which can improve performance by reducing the number of allocations needed as the data structure grows.

		

🧠 Rule of Thumb:-
	Use make(map[K]V, N) when you know the map will hold N items.
It’s a small optimization that can make a big difference in tight loops or large datasets.




	*/
	userName := make([]string,2,5)
	/*
	🧠 Why This Is Useful
- Preallocating capacity avoids multiple memory reallocations during append.
- Great for performance when you know roughly how many items you'll add.
- You can still grow the slice up to 5 elements without triggering a reallocation.

	*/
	userName[0] = "Hi"
	userName = append(userName, "Hello")
	userName = append(userName, "Hey")
	fmt.Println(userName)

	courseRating := make(map[string]float64,3)
	/*
we can add three items here,

without Go having to reallocate memory.

That would only be the case once we add another item here,

then since we only defined a capacity of three here,

Go would have to reallocate memory,

but if we know in advance

that we only plan on adding three items,

at least for the moment,

we could make this a bit more efficient

by using make for creating that map.
*/

courseRating["Go"] = 4.9
courseRating["React"] = 4.8
courseRating["Node"] = 4.9

// using Type Aliases
platformReview:=make(floatMap,3)
platformReview["fast"] = 2.2
platformReview["services"] = 3.1

platformReview.output()

//loop

/*
The range keyword in Go is used for iterating over elements in various data structures like slices, arrays, maps, and channels. It returns the key and value for each element in the data structure.
*/

for key,value := range userName{
	fmt.Println("Key[Index]: ",key)
	fmt.Println("Value: ",value)
}

/*
Additional Notes:
Keys and Values: When using range with a map, the order of iteration over map keys is not guaranteed. It can be different each time you run the program.

Ignored Values: If you only need keys or values, you can ignore the other by using the blank identifier _.


*/

for key,value:=range courseRating{
	fmt.Println(key)
	fmt.Println(value)

}

/*
When iterating over a map using range in Go, the order in which keys and values are accessed is not predictable. This is because maps in Go are inherently unordered collections. The iteration order is based on the internal implementation of the map and can vary between different executions of the same program.

Why is the Order Not Guaranteed?
Hash Table: Go implements maps as hash tables. The keys are hashed and distributed across buckets. The hash function and internal bucket organization influence the iteration order.

Performance Optimization: Maintaining a consistent order would require additional overhead, which could degrade performance. By not guaranteeing an order, Go ensures maps remain efficient for lookups, insertions, and deletions.

Design Choice: Go's design prioritizes simplicity and performance over deterministic iteration order for maps.


Key Points to Remember:
Unordered Nature: Maps are unordered by design. Don't rely on any specific order for iteration.

Consistent Hashing: The hash function's consistency ensures key lookups are efficient but doesn't dictate iteration order.

Use Cases: If you need to maintain order, consider using other data structures like slices along with maps.

By understanding and embracing the unordered nature of maps in Go, you can design your programs to be more robust and efficient, without relying on any specific iteration order.

*/

//structs vs maps
/*
Maps: Dynamic, unordered collections of key-value pairs, excellent for fast lookups and flexible data storage.

Structs: Fixed, ordered collections of fields, ideal for modeling and encapsulating data with type safety and methods.

Understanding when to use maps versus structs can significantly improve the design and performance of your Go applications.
*/
}