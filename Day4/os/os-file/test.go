package main

import (
	"fmt"
	"os"
)

func main(){
	//1.Display Current Working Directory
	cwd,err1:=os.Getwd()
	if(err1!=nil){
		fmt.Println("Error",err1)
		return
	}
	fmt.Println("Current Working Directory", cwd)

//2. List Files in a Directory
dirname:="."
files,err2:=os.ReadDir(dirname)
if(err2!=nil){
		fmt.Println("Error",err2)
}
for _,file:=range files{
	fmt.Println(file.Name())
}

// 3. Create a New Directory
dirname1:="new_directory"
err3:=os.Mkdir(dirname1,0755)

    if err3 != nil {
        fmt.Println("Error:", err3)
        return
    }
    fmt.Printf("Directory '%s' created\n", dirname1)


// 4. Check if a File or Directory Exists
path:="new_directory"
if _,err4:=os.Stat(path);os.IsNotExist(err4){
fmt.Printf("'%s' does not exist\n", path)
}else {
        fmt.Printf("'%s' exists\n", path)
}


//5. Remove a File
filePath := "old_file.txt"
    err5 := os.Remove(filePath)
    if err5 != nil {
        fmt.Println("Error:", err5)
        return
    }
    fmt.Printf("File '%s' removed\n", filePath)
}