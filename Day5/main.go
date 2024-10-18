package main

import (
	"fmt"

	"example.com/main/fileops"
	"github.com/Pallinder/go-randomdata"
)
const accountBalanceFile = "balance.txt"


func main(){
	var accountBalance,err =fileops.GetFloatFromFile(accountBalanceFile)
	if(err!=nil){
		fmt.Println("Error")
	
		fmt.Println(err)
		fmt.Println("---------------")
		// panic("can't continue...")

	}
	fmt.Println("Welcome to the Go Bank")
	fmt.Println("Reach us 24/7",randomdata.PhoneNumber())
	

	for {

options()
var choice int
fmt.Scan(&choice)


switch choice{
case 1: 
fmt.Println("Your Balance is : ",accountBalance)
case 2: 
	fmt.Println("Deposit: ")
	var depositAmount float64
	fmt.Scan(&depositAmount)
	if depositAmount<=0{
		fmt.Println("Invalid Amount")
		continue
	}
	accountBalance+=depositAmount
fmt.Println("Balance updated! New amount : ",accountBalance)
fileops.WriteFloatToFile(accountBalanceFile,accountBalance)
case 3:
	var withdrawlAmount float64
	fmt.Scan(&withdrawlAmount)
	if withdrawlAmount>accountBalance{
		fmt.Println("Invalid Amount")
		continue
	}else if(withdrawlAmount<=0){
		fmt.Println("Invalid Amount")
		continue
	}
	accountBalance-=withdrawlAmount
fmt.Println("Balance updated! New amount : ",accountBalance)
fileops.WriteFloatToFile(accountBalanceFile,accountBalance)
default:
fmt.Println("GoodBye")
fmt.Print("Thanks for choosing our bank")
	return
}
}
}

