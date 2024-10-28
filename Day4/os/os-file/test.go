package main

import (
	"fmt"
	"os"

	// In Go, you can use the reflect package to determine the type of a variable.
	"reflect"
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
// Use reflect.TypeOf from the reflect package to determine the type of a variable.
//When you use os.ReadDir, it returns a slice of fs.DirEntry, representing each file or directory entry in the specified directory.
fmt.Println(reflect.TypeOf(files))
/*
Notes:
fs generally refers to filesystem-related functionality. The fs package within the io library provides interfaces and abstractions for dealing with filesystem operations.
# Key Interfaces in io/fs:
1.fs.FS: The root interface for accessing a filesystem.

2.fs.File: Represents an open file within a filesystem.

3.fs.DirEntry: Represents a directory entry (like a file or directory) within a filesystem.
*/
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
//In Go, you can dynamically declare a variable within an if statement using :=, which is only valid within the scope of the if block. This approach makes the code concise and clean.
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