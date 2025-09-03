// we can use function themselves as parameter values for other functions

package main

import "fmt"

type tranformFn func(int) int


func main(){
	numbers :=[]int{1,2,3,4}
doubled:= transformNumbers(&numbers,double)
tripled := transformNumbers(&numbers,triple)
fmt.Println(doubled)
fmt.Println(tripled)

////
moreNumbers := []int{4,5,6}
// returning functions as a value
t1 := modiefyFunc(&numbers)
t2 := modiefyFunc(&moreNumbers)


v1 := transformNumbers(&numbers,t1)
v2 := transformNumbers(&moreNumbers,t2)


fmt.Println(v1)
fmt.Println(v2)



}



func transformNumbers(numbers *[]int,tranform tranformFn) []int {

dNumbers := []int{}
	for _,val := range *numbers {
dNumbers = append(dNumbers, tranform(val))
	}
return  dNumbers
}




func double(val int)int{
	return  val*2
}
func triple(val int)int{
	return  val*3
}


// function returning as a values:
func modiefyFunc(numbers *[]int) tranformFn{
	// derefernce to get value
	if(*numbers)[0]==1{
		return  double
	}
	return  triple
}