package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
firstName string
lastName string
birtDate string
createAt time.Time

}
type Admin struct{
	email string
	password string
	User
}

func (u *User)  OutputUserDetails(){
  
  fmt.Println(u.firstName,u.lastName,u.birtDate)
}

func (u *User) ClearUserName(){
  u.firstName=""
  u.lastName=""
  u.birtDate=""
}

func New (firstName,lastName,birtDate string) (*User,error){
  //Validation
  if(firstName==""||lastName==""||birtDate==""){
    return nil, errors.New("first name, last name and birthdate are required")
  }
  return  &User{
     firstName,
    lastName,
  birtDate,
     time.Now(),

  },nil
}

func NewAdmin(email,password string) (Admin){
	return Admin{
		email: email,
		password:password,
		User: User{
			firstName: "Admin",
			lastName: "Admin",
			birtDate: "----",
		},
	}
}
