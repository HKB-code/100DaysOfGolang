package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
)

func main(){
	// 1.Display Current Working Directory
	cwd,err1:=os.Getwd()
	if(err1!=nil){
		fmt.Println("Error",err1)
		return
	}
	fmt.Println("Current Working Directory", cwd)

// 2.List Files in a Directory
dirName:="."
files,err2:=os.ReadDir(dirName)
if(err2!=nil){
	fmt.Println("Error", err2)
}
for _,file := range files{
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
dirName1:="new_directory"
err3:=os.Mkdir(dirName1,0755)

if err3 != nil {
	fmt.Println("Error: ", err3)
	return
}
fmt.Printf("Directory '%s' Created\n", dirName1)


// 4.Check if a file or directory exists
path := "new_directory"
//In Go, you can dynamically declare a variable within an if statement using :=, which is only valid within the scope of the if block. This approach makes the code concise and clean.

 if _,err4:=os.Stat(path);
 os.IsNotExist(err4){
	fmt.Printf("'%s' does not exist\n",path )
 }else{
        fmt.Printf("'%s' exists\n", path)

 }



//  5.Remove a File
filePath:= "old_file.txt"
err5:=os.Remove(filePath)
if  err5 != nil {
	fmt.Println("Error", err5)
}
fmt.Printf("File '%s' removed\n", filePath)



// 6.
file4,err:=os.Open("nonexixtentFile.txt")
	if(err!=nil){
		log.Fatal(err)

	}
	fmt.Println(file4)
	file4.Close()


	/*
🗂 Go io/fs Package — Key Interfaces
1. fs.FS
- Role: Root abstraction for a filesystem.
- Think of it as: The "entry point" or the filesystem handle.
- Key Method:
- Open(name string) (fs.File, error) → Opens a file by path.
- Analogy: Like a database connection object that lets you query tables; here, it lets you query files

	*/
K:=os.DirFS(".")

file1,err:=K.Open("index.html")
if  err!= nil {
	fmt.Println(err)
}

defer file1.Close()

s:= make([]byte,1024)
n,er:= file1.Read(s)
  if er != nil {
        panic(er)
    }
fmt.Println(string(s[:n]))
fmt.Println("Opened file via fs.FS")

/*
2. fs.File
- Role: Represents an open file.
- Think of it as: A file descriptor with read/close capabilities.
- Key Methods (inherited from interfaces):
- Stat() (fs.FileInfo, error) → Metadata about the file.
- Read([]byte) (int, error) → Reads file contents.
- Close() error → Closes the file.
- Analogy: Similar to a DB cursor—you open it, read rows (bytes), then close it.

*/

info,_:= file1.Stat()
 fmt.Printf("File name: %s, Size: %d bytes\n", info.Name(), info.Size())
/*
 3. fs.DirEntry
- Role: Represents a directory entry (file or folder).
- Think of it as: A lightweight metadata view of a file/directory.
- Key Methods:
- Name() string → Entry name.
- IsDir() bool → Whether it’s a directory.
- Type() fs.FileMode → File type bits.
- Info() (fs.FileInfo, error) → Full metadata.
- Analogy: Like a row in a directory listing—quick info without fully opening the file


🔗 How They Fit Together
- fs.FS → Opens files → returns fs.File.
- Directory traversal (fs.ReadDir) → returns a slice of fs.DirEntry.
- fs.DirEntry → can be used to decide whether to open as fs.File.

⚡ Backend Analogy
Imagine building a file explorer API:
- fs.FS = The root filesystem object (like mounting a drive).
- fs.File = The actual file handle when a user clicks/open.
- fs.DirEntry = The list of items shown in a folder view.

*/
// Notes:
/*
📖 ReadFile
- Function: os.ReadFile(name string) ([]byte, error)
- Defined in: os (and also io/ioutil in older Go versions).
- Purpose: Convenience function to read the entire file into memory at once.
- Behavior:
- Opens the file internally.
- Reads all bytes until EOF.
- Closes the file automatically.
- Return: A []byte containing the full file contents.
- Usage Example:
data, err := os.ReadFile("index.html")
if err != nil {
    panic(err)
}
fmt.Println(string(data)) // prints entire file
- Analogy: Like running a query and fetching all rows at once into memory.


📖 Read
- Method: Read(p []byte) (n int, err error)
- Defined in: io.Reader interface (implemented by fs.File, os.File, etc.).
- Purpose: Reads chunks of data into a buffer you provide.
- Behavior:
- Reads up to len(p) bytes.
- Returns how many bytes were actually read (n).
- Returns io.EOF when the file ends.
- Usage Example:
f, _ := os.Open("index.html")
defer f.Close()

buf := make([]byte, 512)
for {
    n, err := f.Read(buf)
    if n > 0 {
        fmt.Print(string(buf[:n]))
    }
    if err == io.EOF {
        break
    }
    if err != nil {
        panic(err)
    }
}
	- Analogy: Like fetching rows in batches from a cursor.

🖥️Key Differences
1.ReadFile
level: High-level helper function
reads:Entire file at once
memory:Loads whole file into memory
file handle:Opens & closes internally
use case:Small/medium files, quick reads

2.Read
level: Low-level method on a file handle
reads:Chunk by chunk into a buffer
memory:More memory-efficient for large files
file handle:You manage open/close manually
use case:Large files, streaming, partial reads

⚡ Backend Analogy
- ReadFile = SELECT * FROM table → load everything into memory immediately.
- Read = FETCH NEXT 100 ROWS → stream results in controlled chunks.

👉 So:
- Use ReadFile when the file is small and you just want all contents quickly.
- Use Read when the file is large, or when you want to process it piece by piece (streaming, parsing, etc.).

*/

/*
Notes:
🧩 Concrete Type in Go
- A concrete type is a real type (struct, pointer to struct, or even a basic type alias) that provides actual implementations of the methods required by an interface.
- An interface is just a contract — it says “any type that has these methods can be treated as this interface.”
- A concrete type is the thing that actually defines those methods

// Interface
type Shape interface {
    Area() float64
}

// Concrete type
type Rectangle struct {
    Width, Height float64
}

// Method implementation
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

👉 So yes:
A concrete type is nothing but a type that defines and implements the methods required by an interface, making it usable wherever that interface is expected.


In Your Filesystem Example
- fs.FS = interface (defines Open(name string) (fs.File, error)).
- os.dirFS = concrete type (returned by os.DirFS, implements Open).
- fs.File = interface (defines Read, Stat, Close).
- *os.File = concrete type (returned by Open, implements those methods)

*/

// Notes:
/*
- os.Stat(path)
- Takes a file path string
- Works only with the real filesystem
- Doesn’t require opening the file
- Always backed by OS syscall (stat)
-Quick metadata lookup

- file.Stat() (fs.File from os.DirFS.Open)
- Called on an opened file handle
- Works with any fs.FS (real, embed, custom)
- Requires file to be opened first
- Behavior depends on FS implementation (DirFS → OS-backed, embed.FS → compile-time metadata)
-Metadata tied to an opened file, or abstract FS
*/
}