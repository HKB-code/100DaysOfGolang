package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/main/note"
	"example.com/main/todo"
)

// The code is functioning correctly, but it violates the DRY principle by duplicating the Save functionality in both Note and Todo.
// To adhere to DRY, we introduced a Savable interface, allowing both Note and Todo to implement a common Save method.

type saver interface{
	Save() error
}
// Embedded  interface
type outputtable interface{
	saver
	Display()
}
func SaveData(data saver)error{
	err:=data.Save()
	if err!=nil{
  fmt.Println("Saving the data failed")
	return err
}
fmt.Println("Saving the data succeeded")
return nil
}
func outputData(data outputtable)error{
	data.Display()
return SaveData(data)

}

func getInputs(text string)string {
fmt.Printf("%v ",text)

	reader:= bufio.NewReader(os.Stdin)
	text,err:=reader.ReadString('\n')
	if  err!= nil {
		return ""
	}
text = strings.TrimSuffix(text,"\n")
text = strings.TrimSuffix(text,"\r")
	return text
}


func getNoteData()(string ,string){
	title :=getInputs("Note title: ")



	content:=getInputs("Note content:")

return  title,content
}
func main() {
	printSomething(1)
	printSomething(1.2)
	title,content:=getNoteData()
	todoText := getInputs("Todo texts: ")
	todo,err:=todo.New(todoText)
	if err != nil {
	fmt.Println(err)
	return 
}
err=outputData(todo)
if err!=nil{
	return
}
userNote,err:=note.New(title,content)
if err != nil {
	fmt.Println(err)
	return 
}
outputData(userNote)

}


// for any type value
func printSomething(value interface{}){

intValue,ok:=value.(int)
if(ok){
		fmt.Println("Integer: ",intValue)
		return
}
stringValue,ok :=value.(string)

	if(ok){
		fmt.Println("String: ",stringValue)
		return
	}
			floatValue,ok:=value.(float64)
	if(ok){
		fmt.Println("Float: ",floatValue)
		return
	}



	// you can also take a look at the type of a value with this special syntax.
// 	switch value.(type){
// 		case int: 
// 		fmt.Println("Integer: " , value)
// 	case float64:
// 		fmt.Println("Float: " , value)
// 	case string:
// 		fmt.Println("String: " , value)
//   default:
// 	fmt.Println(value)
// 	}

}