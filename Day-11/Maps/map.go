package main

import "fmt"

/*
Map:
A map is a built-in data type in Go that associates keys with values. It's similar to dictionaries in other languages.

Characteristics of Maps:
Dynamic Size: Maps can grow and shrink dynamically as you add or remove key-value pairs.

Key-Value Pair: Each element in a map is a key-value pair, where each key must be unique.

Unordered: Maps are unordered collections. The order of iteration over map keys is not specified and can vary.

Flexible Keys: Keys can be of almost any type, except slices, maps, and functions.

Efficient Lookup: Maps provide efficient lookups for keys.
*/

func main(){
	// websites := []string{"Google","https://google.com",}
	// fmt.Println(websites)
websites := map[string]string{"Google": "https://google.com",
"Amazon Web Services": "https://aws.com",

}
fmt.Println(websites)
fmt.Println(websites["Amazon Web Services"])
websites["Linkdin"] = "https://linkdin.com"

// delete 
delete(websites,"Google")
fmt.Println(websites)
}