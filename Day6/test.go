//part-1
// package main

// import (
// 	"fmt"
// 	"time"
// )
// func getuserData(text string)string{
//   fmt.Println(text)
//   var value string
// fmt.Scan(&value)
// return value
// }

// type User struct {
// firstName string
// lastName string
// birtDate string
// createAt time.Time

// }
// //(u User ): called Receiver

// /*
// //uses a copy
// func (u User)  outputUserDetails(){

//   fmt.Println(u.firstName,u.lastName,u.birtDate)
// }

// func (u User) clearUserName(){
//   u.firstName=""
//   u.lastName=""
//   u.birtDate=""
// }
// */
// func (u *User)  outputUserDetails(){

//   fmt.Println(u.firstName,u.lastName,u.birtDate)
// }

// func (u *User) clearUserName(){
//   u.firstName=""
//   u.lastName=""
//   u.birtDate=""
// }

// func main(){
//   userfirstName:=getuserData("Please enter your first name: ")
//   userlastName:=getuserData("Please enter your last name: ")
//   userbirtDate:=getuserData("Please enter your birthdate (MM/DD/YYYY): ")

//   var appUser User
//   // appUser= User{
//   //   firstName: userfirstName,
//   //   lastName: userlastName,
//   //   birtDate: userbirtDate,
//   //   createAt: time.Now(),

//   // }
//   // null value for User Struct
//   // appUser=User{}

//     appUser= User{
//      userfirstName,
//      userlastName,
//     userbirtDate,
//      time.Now(),

//   }

//   appUser.outputUserDetails()

// appUser.clearUserName()
// appUser.outputUserDetails()

// // outputUserDetails(userfirstName,userlastName,userbirtDate)
// // outputUserDetails(&appUser)
// }

// // func outputUserDetails(firstName,lastName,birtDate string){
// //   fmt.Println(firstName,lastName,birtDate)

// // }
// ///////////////
// // func outputUserDetails(u *User){
// //   //(*u).firstName
// //   fmt.Println(u.firstName,u.lastName,u.birtDate)
// // }
// //////////////////////////////////////
// part-2
package main

import (
	"errors"
	"fmt"
	"time"
)
func getuserData(text string)string{
  fmt.Println(text)
  var value string
  /*
  fmt.Scanln() and fmt.Scan() are both part of the Go programming language for reading input, but there are key differences:

fmt.Scanln(): Reads input from the standard input until a newline character is encountered. It consumes the newline character.

fmt.Scan(): Reads input from the standard input and stops reading at whitespace (spaces, tabs, or newlines). It doesn't consume the newline character.

Use fmt.Scanln() when you want to read entire lines or when you know there will be newline characters in your input. Use fmt.Scan() when you want to read input separated by spaces.
  */
fmt.Scanln(&value)
return value
}


type User struct {
firstName string
lastName string
birtDate string
createAt time.Time

}

func (u *User)  outputUserDetails(){
  
  fmt.Println(u.firstName,u.lastName,u.birtDate)
}
/*
A constructor function is simply a function that initializes a new instance of a struct and returns it, often with some fields set to default or computed values.
*/
func (u *User) clearUserName(){
  u.firstName=""
  u.lastName=""
  u.birtDate=""
}
//constructor
/*
Now, by returning a pointer here instead of a value,

we again prevent this value from being copied

because if you would return a regular value,

a copy would be created.

That's not just the case for received arguments

but also for returned values.

And therefore, by returning a pointer, we avoid this*/
// func newUser(firstName,lastName,birtDate string) User{
//   return  User{
//      firstName,
//     lastName,
//   birtDate,
//      time.Now(),

//   }
// }
  /*
  Why Use Pointers?
Avoid Copying: Returning a pointer from the constructor avoids copying the struct, which can be more efficient, especially for large structs.

Modify Original: Using pointers allows methods to modify the original struct, as seen with clearUserName.*/
/*
It returns a pointer to a new User struct, initializing the createAt field with the current time.
*/
func newUser(firstName,lastName,birtDate string) (*User,error){
  //Validation
  if(firstName==""||lastName==""||birtDate==""){
    return nil, errors.New("First name, last name and birthdate are required...")
  }
  return  &User{
     firstName,
    lastName,
  birtDate,
     time.Now(),

  },nil
}


func main(){
  userfirstName:=getuserData("Please enter your first name: ")
  userlastName:=getuserData("Please enter your last name: ")
  userbirtDate:=getuserData("Please enter your birthdate (MM/DD/YYYY): ")
/*Now, when returning a pointer,

you also have to edit this code down here though

because now appUser will not store a user object

instead a pointer to a user object

because that is what newUser returns,*/
  // var appUser User

  var appUser *User

  
  //   appUser= User{
  //    userfirstName,
  //    userlastName,
  //   userbirtDate,
  //    time.Now(),

  // }
  //using constructor function
  appUser,err := newUser(userfirstName,userlastName,userbirtDate)
if(err!=nil){
  fmt.Println(err)
  return
}
  appUser.outputUserDetails()

appUser.clearUserName()
appUser.outputUserDetails()




}

