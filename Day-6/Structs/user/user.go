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
	// compostion
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

/*
Go Struct: Composition vs Embedding – Key Notes

User User   // named field

→ This is Composition (Admin has-a User).
Composition (named field):Access: admin.User.firName, admin.User.OutputString()
Methods NOT promoted → must use .User.Method()

Embedding (for promotion):go

User        // no field name → anonymous field

→ This is Embedding.
Embedding benefits:Fields promoted → admin.firName (direct access)
Methods promoted → admin.OutputString(), admin.ClearUserName() (direct call)

Method receiver rules (works in both cases):Method defined on *User (pointer receiver)
Can call on pointer: appUser.OutputString() ✓
Can call on value: admin.User.OutputString() ✓ → Go auto takes address (&admin.User)

Best choice:Want clear separation → use Composition (User User)
Want inheritance-like behavior → use Embedding (User)


*/