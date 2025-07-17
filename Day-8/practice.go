package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/practice/note"
)

func getInputs(text string)(string){
fmt.Printf("%v ",text)
reader:=bufio.NewReader(os.Stdin)
text,err := reader.ReadString('\n')
if err != nil {
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
	title,content:=getNoteData()
userNote,err:=note.New(title,content)
if err != nil {
	fmt.Println(err)
	return 
}
userNote.Display()
err= userNote.Save()
if err != nil {
	fmt.Println("Saving the note failed")
	return 
}
fmt.Println("Saving the note succeeded!")
}