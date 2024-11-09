package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

//constructor
func New(text string) (Todo, error) {
	if text == "" {
		return Todo{}, errors.New("invalid inputs")
	}
	return Todo{
		Text: text,
	}, nil
}

//Display
func (todo Todo) Display(){
	fmt.Println("Text: ", todo.Text)

}

//Save
func (todo Todo) Save() error{
	fileName:= "todo.json"
	json,err:=json.Marshal(todo)
	if(err!=nil){
		return err
	}
	return os.WriteFile(fileName,json,0644)
}