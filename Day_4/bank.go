package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalance = "balanceTxt.txt"
func writeBalanceFile(balance float64){
	balanceTxt := fmt.Sprint(balance)
	/*[]byte(balanceTxt) converts the string balanceTxt into a UTF-8 encoded byte slice, which is the format os.WriteFile expects for writing data to a file. This ensures compatibility, as Go strings are UTF-8 by default, and the byte slice represents the string's characters in a way that can be written to a file.

*/
	os.WriteFile(accountBalance,[]byte(balanceTxt),0644)
}
/*
In Go, error handling is designed to prevent application crashes, unlike some languages where errors can abruptly terminate the program. For example, the os.ReadFile function doesn't crash the app if a file is not found; instead, it returns an error alongside an empty byte slice ([]byte). This empty slice is converted to an empty string, which, when parsed as a float64, results in a value of 0.0. Unlike languages that rely on try-catch blocks to handle exceptions, Go encourages explicit error checking. Functions are written to return errors as values, allowing developers to handle them gracefully without crashing the application. This approach promotes robust and predictable error management, ensuring the program continues running even when issues like missing files occur.


*/
func getBalanceFile()(float64,error){
	data,err:=os.ReadFile(accountBalance)
	if err!=nil{
		return 1000,errors.New("Failed to find balance file")
	}
	balanceTxt:=string(data)
	balance,err:= strconv.ParseFloat(balanceTxt,64)
	// fmt.Println(balanceTxt)
		if err!=nil{
		return 1000,errors.New("Failed to parse stored balance value.")
	}
	return balance,nil
}

func main() {
var accountBalance,err = getBalanceFile()
if err!=nil{
	fmt.Println("Error")
	fmt.Println(err)
	fmt.Println("-----------")
}

	fmt.Println("Welcome to the Go Bank!")
	for{
fmt.Println("What do you want to do?")
	fmt.Println("1.Check Balance")
	fmt.Println("2.Deposit money")
	fmt.Println("3.Withdraw money")
	fmt.Println("4.Exit")

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
		writeBalanceFile(accountBalance)



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
		writeBalanceFile(accountBalance)


	default: 
           fmt.Println("GoodBye")
	fmt.Println("Thank you for choosing us")

		return
		// break
	}
}

	
}
