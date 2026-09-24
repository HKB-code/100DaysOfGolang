package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// io.Reader


func main() {
	// 1.File
	file,_ := os.Open("index.txt")
	readAndPrint(file)
	file.Close()


// 2.String
s:= strings.NewReader("Hello from string")
readAndPrint(s)

}


func readAndPrint(r io.Reader){
	buf := make([]byte,1024)
	


	for {
n,err:= r.Read(buf)
if n>0{
	fmt.Println(string(buf[:n]))
}
if err == io.EOF{
	fmt.Println(err)
	return 
}
if err!=nil{
	fmt.Println(err)
}
	}
	
}