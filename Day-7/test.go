// custom types and adding method
package main

import "fmt"


type str string

func (text str) log(){
	fmt.Println(text)
}

func main(){
	var ok str = "Hello"
	ok.log()
}