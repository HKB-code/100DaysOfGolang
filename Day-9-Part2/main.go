package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	note "example.com/main/note"
	todo "example.com/main/todo"
)
type saver interface{
	Save() error
}
type outputtable interface{
	saver
	Display()
} 


func SaverData(data saver)error{
err:= data.Save()
if  err!= nil {
	fmt.Println("Saving the data failed")
	return err
}
fmt.Println("Saving the data succeeded")
return nil
}

func outputData(data outputtable)error{
	data.Display()
	return SaverData(data)

}

func process (data outputtable){
	if err:= outputData(data);err!=nil{
		fmt.Printf("Operation failed %v\n", err)
		os.Exit(1)
	}
}

func handleErr(err error){
	if err!=nil{
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}


func getInputs(text string) string {
	fmt.Println(text)
	scanner:=bufio.NewScanner(os.Stdin)
	if scanner.Scan(){
		return  scanner.Text()
	}
	if err:=scanner.Err();err!=nil{
		fmt.Print("Error",err)
	}
	return ""
}

func getMultipleLinesInput(prompt string)string{
	fmt.Println(prompt)
	var lines []string
	scanner:=bufio.NewScanner(os.Stdin)

    for scanner.Scan(){
		line := strings.TrimSpace(scanner.Text())
		if line==""{
			break
		}
		lines = append(lines,scanner.Text() )

		    if err := scanner.Err(); err != nil {
        fmt.Println("Input error:", err)
        return ""
    }
	if len(lines)==0{
		return  ""
	}
	}
	return  strings.Join(lines,"\n")
}


func getNoteData()(string,string){
	title :=getInputs("Note title: ")



	content:=getMultipleLinesInput("Note content:")
	return  title,content
}

func main(){
	title,content:=getNoteData()
	todoText:= getMultipleLinesInput("Enter Todo:")
	// todo define
	todo,err:= todo.New(todoText)
	handleErr(err)
process(todo)

	// note define
	userNote,err:= note.New(title,content)
	handleErr(err)

process(userNote)
}