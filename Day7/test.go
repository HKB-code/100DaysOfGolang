package main

import (
	"fmt"

	"example.com/struct/user"
)
func getuserData(text string)string{
  fmt.Println(text)
  var value string
 
fmt.Scanln(&value)
return value
}




func main(){
  userfirstName:=getuserData("Please enter your first name: ")
  userlastName:=getuserData("Please enter your last name: ")
  userbirtDate:=getuserData("Please enter your birthdate (MM/DD/YYYY): ")


  var appUser*user.User

  

  appUser,err:= user.New(userfirstName,userlastName,userbirtDate)
if(err!=nil){
  fmt.Println(err)
  return
}



  appUser.OutputUserDetails()

appUser.ClearUserName()
appUser.OutputUserDetails()



admin:=user.NewAdmin("email@gmail.com","1234")
admin.OutputUserDetails()
admin.ClearUserName()

}

