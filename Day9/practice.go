package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"practice.com/note"
)

func getUserInput(prompt string)(string) {
fmt.Println(prompt)
  // var value string
	// fmt.Scanln(&value)
  reader:=bufio.NewReader(os.Stdin)
  text,err:= reader.ReadString('\n')
  if(err!=nil){
    return ""

  }
  text = strings.TrimSuffix(text,"\n")
  text = strings.TrimSuffix(text,"\r")

return text


}

func getNoteData()(string,string){
  title:=getUserInput("Note Title:")
  
  content:=getUserInput("Note Content:")
   
  return title,content
}


func main() {

title,content:=getNoteData()

userNote,err:=note.New(title,content)
if(err!=nil){
  fmt.Println(err)
  return
}
// fmt.Println(userNote)
userNote.Display()
}