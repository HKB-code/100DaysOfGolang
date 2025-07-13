package main

import (
	"fmt"

	"example.com/struct/user"
)


func main() {
	userfirstName:= getUserDate("Please enter your first name")
	userlastName:= getUserDate("please enter your last name")
	userbirtDate:= getUserDate("please enter you dob (MM/DD/YYYY)")


var appUser *user.User
appUser,err:= user.NewUser(userfirstName,userlastName,userbirtDate)
if err != nil {
	fmt.Println(err)
	return
}

// appUser = user{
// 	firstName: userfirstName,
// 	lastName: userlastName,
// 	createdAt: time.Now(),
// 	birthdate: userbirtDate,
// }

appUser.OutputString()
appUser.ClearUserName()
appUser.OutputString()





admin := user.AdminNew("email@gmail.com","1234")
admin.User.OutputString()


}

func getUserDate(promptText string) string {
	fmt.Println(promptText)
	var value string
	fmt.Scanln(&value)
	return  value
}

