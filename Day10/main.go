// The code is functioning correctly, but it violates the DRY principle by duplicating the Save functionality in both Note and Todo.
// To adhere to DRY, we introduced a Savable interface, allowing both Note and Todo to implement a common Save method.
/*
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/main/notes"
	"example.com/main/todo"


)

func getUserInput(prompt string) string {
	fmt.Println(prompt)
	reader:=bufio.NewReader(os.Stdin)
	text,err:=reader.ReadString('\n')
	if(err!=nil){
		return""
	}
	text = strings.TrimSuffix(text,"\n")
	text = strings.TrimSuffix(text,"\r")
	return text

}

func getNotsData()(string, string){
	title:=getUserInput("Note Title: ")
	content:= getUserInput("Note Content: ")
	return title, content
}

func main(){
	title,content:=getNotsData()
	userNote,err:= notes.New(title,content)
	if(err!=nil){
		fmt.Println(err)
		return
	}
	userNote.Display()

	err = userNote.Save()
if err!=nil{
  fmt.Println("Saving the note failed")
}
fmt.Println("Saving the note succeeded")

text:= getUserInput("Text: ")
todoUser,err:=todo.New(text)
	if(err!=nil){
		fmt.Println(err)
		return
	}

todoUser.Display()
err=todoUser.Save()
if err!=nil{
  fmt.Println("Saving the note failed")
}
fmt.Println("Saving the note succeeded")
}
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/main/notes"
	"example.com/main/todo"
)

//interface
/*
Just like a public parking lot can be used by any car that fits, an interface in Go can be used by any type that implements the required methods. It doesn't matter what brand or type of car (or struct) it is, as long as it follows the parking rules (methods), it can park (implement the interface) there. 🚗🔧*/
type Saver interface{
	Save() error
	Display()
/*
because this interface, the saver interface,

simply says that a valid value

must have a save method, which takes no arguments

and returns an error.
*/
}
// type Displayer interface{
// 	Display()
// }
//Embedded Interface
type Outputtable interface{
	Saver
	Display()
}


func getUserInput(prompt string) string {
	fmt.Println(prompt)
	reader:=bufio.NewReader(os.Stdin)
	text,err:=reader.ReadString('\n')
	if(err!=nil){
		return""
	}
	text = strings.TrimSuffix(text,"\n") 
	text = strings.TrimSuffix(text,"\r") 
	return text

}

func getNotsData()(string, string){
	title:=getUserInput("Note Title: ")
	content:= getUserInput("Note Content: ")
	return title, content
}


func Savedata(data Saver) error{
err:=data.Save()
if err!=nil{
  fmt.Println("Saving the note failed")
	return err
}
fmt.Println("Saving the note succeeded")
return nil
}
func outputData(data Outputtable)error{
	data.Display()
return Savedata(data)

}

//Any type of value
func printSomething(value interface{}){
	switch value.(type){
	case int: 
		fmt.Println("Integer: " , value)
	case float64:
		fmt.Println("Float: " , value)
	case string:
		fmt.Println("String: " , value)
  default:
	fmt.Println(value)

	}


}
//another way
func anotherPrintSomething(value interface{}){
	intValue,ok:=value.(int)
	if(ok){
		fmt.Println("Integer: ",intValue)
		return
	}
		stringValue,ok:=value.(string)
	if(ok){
		fmt.Println("String: ",stringValue)
		return
	}
		floatValue,ok:=value.(float64)
	if(ok){
		fmt.Println("Float: ",floatValue)
		return
	}
}

func main(){
	printSomething(1)
	printSomething(1.5)
	printSomething("hello")
	fmt.Println("////////////")
	anotherPrintSomething(1)
	anotherPrintSomething(1.5)
anotherPrintSomething("Hello")
	title,content:=getNotsData()
	userNote,err:= notes.New(title,content)
	if(err!=nil){
		fmt.Println(err)
		return
	}
	// userNote.Display()
err =outputData(userNote)
if err!=nil{
	return
}
printSomething(userNote)
// 	err = userNote.Save()
// if err!=nil{
//   fmt.Println("Saving the note failed")
// }
// fmt.Println("Saving the note succeeded")
// err = Savedata(userNote)
// if err!=nil{
// 	return
// }

text:= getUserInput("Text: ")
todoUser,err:=todo.New(text)
	if(err!=nil){
		fmt.Println(err)
		return
	}

// todoUser.Display()
outputData(todoUser)
// err=todoUser.Save()
// if err!=nil{
//   fmt.Println("Saving the note failed")
// }
// fmt.Println("Saving the note succeeded")

// err = Savedata(todoUser)
// if err!=nil{
// 	return
// }
}

