package todo

import (
	"errors"
	"fmt"

	storage "example.com/main/storage"
)

type ToDo struct {
	Text string 
}
func (todo *ToDo)Display(){
fmt.Println((*todo).Text)
}

func (todo *ToDo)Save()error{
	fileName:= "my_todo.json"
return storage.SaveToJSON(todo,fileName)
}

func New(text string)(*ToDo,error){
	if(text==""){
		return  nil,errors.New("invalid input")
	}
	return  &ToDo{
	Text: text,

	},nil
}