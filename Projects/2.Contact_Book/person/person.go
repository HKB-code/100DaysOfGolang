package person

import (
	"errors"
	"fmt"

	"example.com/contact_book/address"
)

type Person struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Address *address.Address `json:"address"`
	Id int `json:"id"`
}

func New(name, email, phone string, add *address.Address,id int)( *Person,error ){
	if (name == "" || email == "" || phone == "" || add==nil||id==0 ){
		return nil, errors.New("please enter all fields")
	}
	return  &Person{
		Name: name,
		Email: email,
		Phone: phone,
		Address: add,
		Id:id,
	},nil
}
func (p *Person) String() string {
    return fmt.Sprintf(
        "ID: %d\nName: %s\nEmail: %s\nPhone: %s\nAddress: %s, %s, %s, %s - %s",
        p.Id,
        p.Name,
        p.Email,
        p.Phone,
        p.Address.Street,
        p.Address.City,
        p.Address.State,
        p.Address.Country,
        p.Address.ZipCode,
    )
}