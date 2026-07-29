package address

import "errors"

type Address struct {
	Street  string
	City    string
	State   string
	Country string
	ZipCode string
}

func New(street, city, state, country, zipcode string) (*Address, error) {
	if street == "" || city == "" || state == "" || country == "" || zipcode == "" {
		return nil, errors.New("please enter all fields")
	}
	return &Address{
		Street:  street,
		City:    city,
		State:   state,
		Country: country,
		ZipCode: zipcode,
	}, nil
}