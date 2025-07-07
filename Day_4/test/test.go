package main

import (
	"errors"
	"fmt"
	"os"
)
const accountFile = "account.txt"

func writeToFile(ebt,profit,ratio float64){
		results:=fmt.Sprintf("EBT: %.2f\nProfit: %.2f\nRatio: %.2f",ebt,profit,ratio)
		os.WriteFile(accountFile,[]byte(results),0644)

}

func getUserInput(text string) (float64,error) {
	var value float64
	fmt.Println(text)
	fmt.Scan(&value)
	if value<=0{
		return 0,errors.New("Please enter a valid number")
	}
	return value,nil
}

func main() {
revenue,err1:=getUserInput("Revenue: ")
expenses,err2:= getUserInput("Expenses: ")
taxRate,err3:= getUserInput("Tax Rate: ")

if err1!=nil||err2!=nil||err3!=nil{
	fmt.Println(err1)
	return
}

ebt,profit,ratio := calc(revenue,expenses,taxRate)
fmt.Printf("ebt: %.2f\n",ebt)
fmt.Printf("Profit: %.2f\n",profit)
fmt.Printf("Ratio: %.2f\n",ratio)
}

func calc(revenue,expenses,taxRate float64)(ebt,profit,ratio float64){
	ebt = revenue-expenses
	profit = ebt*(1-taxRate/100)
	ratio = ebt/profit
writeToFile(ebt,profit,ratio)

	return
}