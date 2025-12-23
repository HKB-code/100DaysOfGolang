/*

Exercise: Maps
Implement WordCount. It should return a map of the counts of each “word” in the string s. The wc.Test function runs a test suite against the provided function and prints success or failure.
*/

package main

import (
	"strings"

	"golang.org/x/tour/wc"
)


func WordCount(s string)map[string]int{

	a:= strings.Fields(s)
	z:= map[string]int{}
	for _ ,v:= range a{
z[v]++

	}
	return  z
}

func main(){
wc.Test(WordCount)
}
