package main

import "fmt"

func Add(a,b int) int{
	return a+b
}

func main(){
var a int = 1000  //explicitly declared as int
var b float64 = 1.0 //explicitly declared as float64
c:=5

res := int(float64(a) +b )
final := Add(res ,c )
	fmt.Println(final)

}