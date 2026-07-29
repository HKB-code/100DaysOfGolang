package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"example.com/contact_book/addressbook"
	"example.com/contact_book/menu"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	book := addressbook.New()
	book.LoadData()
	fmt.Println("Welcome To the Contact Book")
	for{
		menu.Options()
		choice,err:= getInts("Enter Choice: ")
		if  err!= nil {
			fmt.Println("please enter a valid number..")
			continue
		}
		if choice>7||choice<1{
			fmt.Println("please enter a number between 1 and 7")
			continue
		}
		switch choice{
		case 1:
			name:= getStrings("Enter Name: ")
			email:= getStrings("Enter Email: ")
			phone := getStrings("Enter Phone: ")
			street:=getStrings("Enter Street: ")
			city:=getStrings("Enter City: ")
			state:=getStrings("Enter State: ")
			country:= getStrings("Enter Country: ")
			zipcode:= getStrings("Enter Zipcode: ")


			err:=book.Add(name,email,phone,street,city,state,country,zipcode)
			if  err!= nil {
				fmt.Println(err)
				continue
			}
		case 2:
			id,err:= getInts("Enter the id: ")
			if  err!= nil {
				fmt.Println("Please enter a valid number")
				continue
			}
			err=book.Delete(id)
			if  err!= nil {
				fmt.Println(err)
				continue
			}
		case 3:
			name:= getStrings("Enter Name: ")
			
			phone := getStrings("Enter Phone: ")
			p1,err:=book.Search(name,phone)
			if  err!= nil {
				fmt.Println(err)
				continue
			}
			z:=p1.String()
			fmt.Println(z)
		case 4:
			contacts := book.Sort()
			fmt.Println("Here is the sorted list: ")
			for i:=range contacts{
				z:=contacts[i].String()
				fmt.Println(z)
				fmt.Println("--------------------")
			}
        case 5: 
			id,err:= getInts("Enter the id: ")
			if  err!= nil {
				fmt.Println("Please enter a valid number")
				continue
			}
			
			email:= getStrings("Enter Email: ")
			phone := getStrings("Enter Phone: ")
		err=book.Update(id,phone,email)
			if  err!= nil {
				fmt.Println(err)
				continue
			}

		case 6:
			err:=book.SaveToJson()
			if  err!= nil {
				fmt.Println(err)
				continue
			}
		case 7:
			fmt.Println("GoodBye")
			return
		default:
			fmt.Println("please enter a valid choice")
			continue
		}
	}

}


func getStrings(text string)string{
	fmt.Println(text)
if scanner.Scan(){
	return  scanner.Text()
}
if err:= scanner.Err();err!=nil{
	fmt.Println("input err",err)
}

	
return  ""
}


func getInts(text string)(int,error){
	fmt.Println(text)
	if !scanner.Scan(){
		return  0,scanner.Err()
	}
	d := scanner.Text()

	value,err:=strconv.Atoi(d)
	if  err!= nil {
		return  0,err
	}
	return  value,nil

}
