package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)
func WriteValueFile(value float64,fileName string){
	valueTxt := fmt.Sprint(value)
	/*[]byte(balanceTxt) converts the string balanceTxt into a UTF-8 encoded byte slice, which is the format os.WriteFile expects for writing data to a file. This ensures compatibility, as Go strings are UTF-8 by default, and the byte slice represents the string's characters in a way that can be written to a file.

*/
	os.WriteFile(fileName,[]byte(valueTxt),0644)
}
/*
In Go, error handling is designed to prevent application crashes, unlike some languages where errors can abruptly terminate the program. For example, the os.ReadFile function doesn't crash the app if a file is not found; instead, it returns an error alongside an empty byte slice ([]byte). This empty slice is converted to an empty string, which, when parsed as a float64, results in a value of 0.0. Unlike languages that rely on try-catch blocks to handle exceptions, Go encourages explicit error checking. Functions are written to return errors as values, allowing developers to handle them gracefully without crashing the application. This approach promotes robust and predictable error management, ensuring the program continues running even when issues like missing files occur.


*/
func GetFloatFromFile(fileName string)(float64,error){
	data,err:=os.ReadFile(fileName)
	if err!=nil{
		return 1000,errors.New("Failed to find file")
	}
	valueTxt:=string(data)
	value,err:= strconv.ParseFloat(valueTxt,64)
	// fmt.Println(balanceTxt)
		if err!=nil{
		return 1000,errors.New("Failed to parse stored value.")
	}
	return value,nil
}

/*
Multiple Packages in Same Folder: Not allowed in Go. Each package must be in its own subfolder.
Package Subfolder: Create a subfolder for each package, named after the package (e.g., fileops for a package named fileops).
File Naming: Files within the package folder can have any name; they don’t need to match the package name.
Main Package: Typically resides in the root or its own subfolder (e.g., main). Each main package represents an executable program.


*/