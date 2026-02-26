// Package: 1. Splitting code across multiple files.
// 2. Splitting File across multiple Packages.
// 3. Importing & using custom Packages.

package main

import (
	"fmt"
	"time"

	transactionhistory "example.com/main/TransactionHistory"
	fileops "example.com/main/fileOps"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balanceTxt.txt"


func main() {
var accountBalance,err =fileops.GetFloatFromFile(accountBalanceFile)
if err!=nil{
	fmt.Println("Error")
	fmt.Println(err)
	fmt.Println("-----------")
}
err=transactionhistory.WriteTOFile(fmt.Sprint("Default Balance: ",accountBalance))
if  err!= nil {
	fmt.Println(err)
}

	fmt.Println("Welcome to the Go Bank!")
	fmt.Println("Reach us 24/7",randomdata.PhoneNumber())
	for{

     options()
	var choice int 
	fmt.Println("Your Choice")
	fmt.Scan(&choice)
  
	switch choice{
	case 1: 
		fmt.Println("Your Balance is: ",accountBalance)
 
    case 2:
        fmt.Println("Your deposits: ")
		var depositAmpont float64
		fmt.Scan(&depositAmpont)
		if depositAmpont<=0{
			fmt.Println("Invalid deposit")
			// return
			continue
	}
	    accountBalance+=depositAmpont
		fmt.Println("Balance updates: ", accountBalance)
		fileops.WriteValueFile(accountBalance,accountBalanceFile)
		tx:= fmt.Sprintf("Deposite | Amount: %.2f | Balance: %.2f | Time: %s",depositAmpont,accountBalance,time.Now().Format(time.RFC3339))
		err:=transactionhistory.WriteTOFile(tx)
		if  err!= nil {
			fmt.Println(err)
		}




		case 3: 
		fmt.Println("Withdrawl amount")
		var withdrawlAmount float64
		fmt.Scan(&withdrawlAmount)
		if withdrawlAmount<=0{
			fmt.Println("Invalid withdrawal")
			// return
			continue;
		}
		if withdrawlAmount>accountBalance{
			fmt.Println("Invalid withdrawal")
			// return
			continue
		}
		accountBalance -=withdrawlAmount
		fmt.Println("Balance updates: ", accountBalance)
		fileops.WriteValueFile(accountBalance,accountBalanceFile)
        	tx:= fmt.Sprintf("Withdrawal | Amount: %.2f | Balance: %.2f | Time: %s",withdrawlAmount,accountBalance,time.Now().Format(time.RFC3339))
		err:=transactionhistory.WriteTOFile(tx)
		if  err!= nil {
			fmt.Println(err)
		}

	case 4:
		err:=transactionhistory.ReadFile()
		if  err!= nil {
			fmt.Println(err)
		}


	default: 
           fmt.Println("GoodBye")
	fmt.Println("Thank you for choosing us")

		return
		// break
	}
}

	
}


