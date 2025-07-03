//  //Test code
// package main

// import "fmt"

// func main(){
// 	fmt.Print(1+2+3)
// 	fmt.Println("Hello")
// 	name:="Jhon"
// 	age:=22
// 	fmt.Print("His age is", age," Years")
// 	n,_ := fmt.Print("Name: ", name, " Age: ", age)
//  fmt.Println("Number of bytes written: ",n)
// const n1,dept  = "peter", "cs"
// fmt.Print(n1,"\n",dept)
// }

//  // Investment Calculator
// package main

// import (
// 	"fmt"
// 	"math"
// )

// const inflamation  = 2.5

// func main(){
// var investmentAmount, expectedReturnRate,years float64
// outputText("Invested Amount: ")
// fmt.Scan(&investmentAmount)

// outputText("Expected Returs Rate: ")
// fmt.Scan(&expectedReturnRate)

// outputText("Years: ")
// fmt.Scan(&years)

// var futureValue,futureRate = calculatedFutureValue(investmentAmount,expectedReturnRate,years)

// formatedFutureValue := fmt.Sprintf("Future Value: %.2f\n",futureValue)
// formatedFutureRate :=fmt.Sprintf("Future Value (adjusted for inflation): %.2f\n",futureRate)
// fmt.Println(formatedFutureValue,formatedFutureRate)
// }

// func outputText(text string){
// 	fmt.Println(text)
// }

// func calculatedFutureValue(investmentAmount,expectedReturnRate,years float64)(fv, rfv float64){
// 	fv = investmentAmount*math.Pow(1+expectedReturnRate/100,years)
// 	rfv = fv/math.Pow(1+inflamation/100,years)
// 	return
// }

// update code
package main

import (
	"fmt"
	"math"
) 
const inflamationRate = 2.5


func updateValue(text string) float64{
	var value float64
	fmt.Println(text)
	fmt.Scan(&value)
	return value
}

func main(){
	 investmentAmount :=updateValue("Invested Amount: ")
	 expectedAmount:= updateValue("Expected Return Rate: ")
	 years:= updateValue("Years: ")

	 	futureValue,realFutureValue  := calculateFutureValue(investmentAmount,expectedAmount,years)
	
fmt.Println(futureValue)
fmt.Println(realFutureValue)

formattedFv := fmt.Sprintf("Future Value: %.2f\n", futureValue)
fmt.Println(formattedFv)
formattedrFv:=fmt.Sprintf("Real Future Value: %.2f\n",realFutureValue)
fmt.Println(formattedrFv)
}


func calculateFutureValue(investmentAmount,expectedAmount,years float64)(fv float64, rfv float64){
	fv = investmentAmount*math.Pow(1+expectedAmount/100,years)
	rfv = fv/math.Pow(1+inflamationRate/100,years)
	return
}