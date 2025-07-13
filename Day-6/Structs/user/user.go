package user

import (
	"errors"
	"fmt"
	"time"
)

	
type User struct{
	firstName string
	lastName string
	age int
	createdAt time.Time
	birthdate string

}

type Admin struct{
	Email string
	Password string
	User User
}

func  AdminNew(email,password string) Admin{
return  Admin {
	Email: email,
	Password: password,
	User: User{
		firstName: "Admin",
		lastName: "Admin",
		birthdate: "-----",
	},
}
}


func (u *User)OutputString(){
	fmt.Println(u.firstName,u.lastName,u.birthdate)

}
func (u *User) ClearUserName(){
	u.firstName = ""
	u.lastName = ""
	u.birthdate = ""
}
func NewUser(firstName,lastName,birthdate string) (*User,error){
	if(firstName==""||lastName==""||birthdate==""){
		return  nil,errors.New("first name . last name and birtdate are required")
	}
	return &User{
		firstName:firstName ,
		lastName: lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	},nil
}
