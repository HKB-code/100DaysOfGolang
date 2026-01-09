package cmdmanager

import "fmt"

type CMDmanager struct {
}

func New()*CMDmanager{
	return & CMDmanager{}
}

func (cmd CMDmanager)  ReadFiles() ([]string, error) {
	fmt.Println("Please enter your prices. Confirm every price with ENTER.")
	var prices []string

	for{
		var price string
		fmt.Println("Price: ")
		fmt.Scan(&price)
		if price == "0"{
			break
		}
		prices = append(prices, price)
	}
	return  prices,nil
}

func (cmd CMDmanager) WriteResult(data any) error {
	fmt.Println(data)
	return  nil
}