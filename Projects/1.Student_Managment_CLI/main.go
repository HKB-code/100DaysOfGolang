package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"example.com/student_managment_cli/manager"
	"example.com/student_managment_cli/menu"
)
var scanner= bufio.NewScanner(os.Stdin)
func main() {
	mgr := manager.New()
	fmt.Println("Welcome To The Student Managment CLI")
	for{
		menu.Options()
	choice,err:=getInt("Enter Choice: ")
		if  err!= nil {
		fmt.Println("please enter a valid number")
		continue
	}
	if choice<1||choice>6{
		fmt.Println("Please enter a numbetween 1 and 6.")
		continue
	}

	
		switch choice{
		case 1:
			name:= getString("Enter Student Name: ")
			grade := getString("Enter Student Grade: ")
			age,err := getInt("Enter Student Age: ")
			if err!=nil{
				fmt.Println("Please enter a valid number")
				
				continue
			}
		
			err= mgr.Add(name,grade,age)
			if err!=nil{
				fmt.Println(err)
				continue
			}

		case 2:
			id,err:= getInt("Enter Student Id: ")
				if err!=nil{
					
					fmt.Println("please enter a valid number")
				continue
			}

			err = mgr.Delete(id)
			if err!=nil{
				fmt.Println(err)
				continue
			}
		case 3: 
		   id,err:= getInt("Enter Student Id: ")
				if err!=nil{
					
					fmt.Println("please enter a valid number")
				continue
			}
			age,err:= getInt("Enter Student Age: ")
				if err!=nil{
					
					fmt.Println("please enter a valid number")
				continue
			}
			grade := getString("Enter Student Grade: ")

		    err=mgr.Update(id,age,grade)
			if err!=nil{
				fmt.Println(err)
				continue
			}

        case 4:
			id,err:= getInt("Enter Student Id: ")
				if err!=nil{
					fmt.Println("please enter a valid number")
				
				continue
			}
			student,err:= mgr.Search(id)
			if err!=nil{
				fmt.Println(err)
				continue
			}
			fmt.Printf("Name: %s | Age: %d | Grade: %s | Id: %d",student.Name,student.Age,student.Grade,student.Id)
		
		case 5: 

		mgr.List()
		case 6:
			fmt.Println("GoodBye")
			return
		default:
			fmt.Println("please enter a valid choice")
			continue
		}

	}

}

func getString(text string)string{
	
	fmt.Print(text)
	
	if scanner.Scan(){
		return  scanner.Text()
	}
	if err:=scanner.Err();err!=nil{
		fmt.Println("input error", err)
	}

	return ""
}

func getInt(text string)(int,error){
	
fmt.Print(text)
if !scanner.Scan(){
	return  0 , scanner.Err()
}
d := scanner.Text() 

value,err:= strconv.Atoi(d)
if  err!= nil {
	return 0,err
}
return value,nil
	
}
