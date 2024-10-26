package main

import (
	"fmt"
	"time"
)
func getuserData(text string)string{
  fmt.Println(text)
  var value string
fmt.Scan(&value)
return value
}


type User struct {
firstName string
lastName string
birtDate string
createAt time.Time
}



func main(){
  firstName:=getuserData("Please enter your first name: ")
  lastName:=getuserData("Please enter your last name: ")
  birtDate:=getuserData("Please enter your birthdate (MM/DD/YYYY): ")







outputUserDetails(firstName,lastName,birtDate)
}

func outputUserDetails(firstName,lastName,birtDate string){
  fmt.Println(firstName,lastName,birtDate)

}