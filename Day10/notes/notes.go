package notes

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title      string   `json:"title"`
	Content    string  `json:"content"`
	Created_At time.Time `json:"created_at"`
}

//constructor

func New(title, content string)(Note,error){
	if title==""|| content==""{
		return Note{}, errors.New("invalid inputs")
	}
	return Note{
		Title: title ,
		Content: content,
		Created_At: time.Now(),
	},nil
}

//Display
func (note Note) Display(){
	fmt.Printf("Your note titled %v has the following content:\n\n%v\n",note.Title,note.Content)

}

//Save
func (note Note) Save() error{
	fileName:=strings.ReplaceAll(note.Title," ","_")
	fileName  = strings.ToLower(fileName) + ".json"
	json,err:=json.Marshal(note)
	if(err!=nil){
		return err
	}
	return os.WriteFile(fileName,json,0644)
}