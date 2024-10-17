package main

import (
	"errors"
	"fmt"
	"os"
)
const accountFile = "balance.txt"
func writeToFile(ebt,profit,ratio float64){
	results:=fmt.Sprintf("EBT: %.2f\nProfit: %.2f\nRatio: %.2f",ebt,profit,ratio)
  
os.WriteFile(accountFile,[]byte(results),0644)
}

func getUserInputs(text string )(float64,error ){
	var value float64
	
	fmt.Println(text)
	fmt.Scan(&value)
	if value<0{
		return 0, errors.New("value must be a positive number")
	}
	
		
	return value, nil
}
func calc(revenue,expenses,taxRate float64)(ebt float64,profit float64,ratio float64){
ebt  = revenue-expenses
	profit = ebt*(1-taxRate/100)
	ratio = ebt/profit
writeToFile(ebt,profit,ratio)
	return
}


func main(){
revenue,err1:= getUserInputs("Revenue: ")
// if err!=nil{
// 	return
// }
expenses,err2:=getUserInputs("Expenses: ")
// if err!=nil{
// 	return
// }
taxRate,err3:= getUserInputs("Tax Rate: ")
if err1!=nil||err2!=nil||err3!=nil{
	fmt.Println(err1)
	return
}

ebt,profit,ratio:= calc(revenue,expenses,taxRate)


fmt.Printf("ebt: %.2f\n",ebt )
fmt.Printf("profit: %.2f\n",profit )
fmt.Printf("ratio: %.2f\n",ratio )


}