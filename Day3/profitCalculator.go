package main

import "fmt"

func main(){
revenue :=getUserInput("Revenue: ")
expenses :=getUserInput("Expenses: ")
taxRate:=getUserInput("Tax Rate: ")

ebt,profit,ratio := calc(revenue,expenses,taxRate)
fmt.Printf("ebt: %.2f\n",ebt)
fmt.Printf("Profit: %.2f\n",profit)
fmt.Printf("Ratio: %.2f\n",ratio)


}
func getUserInput(text string)(float64){
	var value float64
	fmt.Println(text)
	fmt.Scan(&value)
	return value
}
func calc(revenue,expenses,taxRate float64)(ebt float64,profit float64, ratio float64){
	ebt = revenue-expenses
	profit = ebt*(1-taxRate/100)
	ratio = ebt/profit
	return
}