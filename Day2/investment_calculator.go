/*
# Test-Code
package main

import "fmt"
func main(){
	fmt.Println("Hello")
	fmt.Print(1+2+3)
	name:="Harshil"
	age:=22;
	fmt.Println("His age is",age,"Years")
	n,_ := fmt.Print("Name: ",name,"Age: ",age)
	fmt.Println("Number of bytes written:",n)
	const n1,dept = "Peter","Cs"
	fmt.Print(name,"\n",dept)
}*/

// package main

// import (
// 	"fmt"
// 	"math"
// )
// const inflamationRate = 2.5
// func main(){
// var investmentAmount,expectedReturnRate,years float64
// outputText("Investment Amount:")
// fmt.Scan(&investmentAmount)

// outputText("Expected Return Rate: ")
// fmt.Scan(&expectedReturnRate)

// outputText("Years: ")
// fmt.Scan(&years)

// var futureValue,futureRate = calculateFutureValue(investmentAmount,expectedReturnRate,years)

// fmt.Println(futureValue)
// fmt.Println(futureRate)

// formattedFv := fmt.Sprintf("Future Value: %.2f\n",futureValue)
// formattedFr :=fmt.Sprintf("Future value (adjusted for inflation): %.2f\n",futureRate)
// fmt.Println(formattedFv,formattedFr)
// }

// func outputText(text string){
// 	fmt.Print(text)
// }

// func calculateFutureValue(investmentAmount,expectedReturnRate,years float64) (fv float64, rfv float64){
// 	fv = investmentAmount*math.Pow(1+expectedReturnRate/100,years)
// 	rfv = fv/math.Pow(1+inflamationRate/100,years)
// 	return
// }

// update code,
package main

import (
	"fmt"
	"math"
)

const inflamationRate = 2.5
func updateValue(text string)(float64){
	var value float64
	fmt.Println(text)
	fmt.Scan(&value)
	return value
}
func main(){
	investmentAmount:=updateValue("Invested Amount: ")
	expectedAmount :=updateValue("Ecpected Return Rate")
	years :=updateValue("Years: ")

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