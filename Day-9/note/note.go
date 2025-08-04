package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("invalid inputs")
	}
	return  Note{
		Title: title,
	Content: content,
	},nil
}

func (note Note)Save()error{
	fileName := strings.ToLower(strings.ReplaceAll(note.Title," ","_")) + ".json"
	json,err := json.Marshal(note)
	if  err!= nil {
		return  err
	}
	
return  os.WriteFile(fileName,json,0644)
}

func (note Note) Display(){
	fmt.Printf("Your note titled %v has the following content:\n%v\n",note.Title,note.Content)

}