package note

import (
	"errors"
	"fmt"
	"strings"
	"time"

	storage "example.com/main/storage"
)

type Note struct {
	Title     string
	Content   string
	CreatedAt time.Time
}
func (note *Note)Display(){
	fmt.Printf("Your note titled %v has the following content:\n%v\n",note.Title,note.Content)
}

func (note *Note)Save()error{
	fileName:= strings.ToLower(strings.ReplaceAll(note.Title," ","_"))+".json"
return storage.SaveToJSON(note,fileName)
}

func New(title,content string)(*Note,error){
	if(title==""||content==""){
		return  nil,errors.New("invalid input")
	}
	return  &Note{
		Title: title,
		Content: content,
		CreatedAt: time.Now(),

	},nil
}