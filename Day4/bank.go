package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)
const accountBalanceFile = "balance.txt"

func getBalanceFromFile() (float64,error){
	
	/*If it does not found balance.text, it going to give 0,Instead in Go, functions are written such

and your functions should be written such

that errors don't crash the application.

Instead, read file for example,

will simply return an empty byte collection

if it fails to find that file,

which is then converted to an empty string,

which is converted to the value zero

when parsing it as a float.*/
/*But that's not all Go does.

Instead, for operations that could fail,

it is quite common that functions return such an error thing

as a second value.

And that's the value we can accept and store

as a second value here when calling that function*/
	data,err:=os.ReadFile(accountBalanceFile)
	/*Now nil is a special value in Go.

It basically stands for the absence of a useful value*/
if err!=nil{
	return 1000, errors.New("failed to find balance file")
}

	balanceText :=string(data)
  balance,err:= strconv.ParseFloat(balanceText,64)
	if(err!=nil){
		return 1000, errors.New("failed to parsed stored balance value")
	}
	return balance,nil
}


func writeBalanceToFile(balance float64){
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile,[]byte(balanceText),0644)
 
}

func main(){
	var accountBalance,err =getBalanceFromFile()
	if(err!=nil){
		fmt.Println("Error")
		//you can refer to the error value like this and it will automatically output the attached message.
		fmt.Println(err)
		fmt.Println("---------------")
		// panic("can't continue...")

	}
	fmt.Println("Welcome to the Go Bank")
	

	for {
	fmt.Println("What do you want to do?")
	fmt.Println("1.Check Balance")
	fmt.Println("2.Deposit Money")
	fmt.Println("3.Withdraw Money")
	fmt.Println("4.Exit")

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
writeBalanceToFile(accountBalance)
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
writeBalanceToFile(accountBalance)
default:
fmt.Println("GoodBye")
fmt.Print("Thanks for choosing our bank")
	return
}
}
}