package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

func main(){
	// 6. Get system information
	fmt.Println("Go Version:", runtime.Version())
    fmt.Println("OS:", runtime.GOOS)
    fmt.Println("Architecture:", runtime.GOARCH)


	
// 7. Execute a Shell Command
/*For Windows Systems:

cmd: The command-line interpreter for Windows.

/C: Flag to carry out the command specified by the string and then terminate.

echo Hello, World!: The command to run.*/
/*Using cmd with /C in Windows is essential because echo is a built-in command in the Windows Command Prompt, and cmd is the command-line interpreter for executing it.

Breakdown:
cmd:

The command-line interpreter for Windows, used to run commands in a command prompt.

/C:

Tells cmd to run the specified command and then terminate.

Without /C, the command prompt would stay open or the command might not execute properly.

echo Hello, World!:

The actual command to be executed.

echo is a built-in command in the Windows Command Prompt that outputs the text following it.

Why These Steps Are Needed:
Execute Built-in Commands: echo and similar commands are built-in commands of the command interpreter (cmd). They are not standalone executable files.

Ensure Proper Execution and Termination: Using cmd /C ensures that the command runs in the command prompt environment and the environment is closed after execution.*/
		cmd := exec.Command("cmd","/C", "echo Hello, World!")
    output, err := cmd.Output()
    if(err!=nil){
        fmt.Println("Error:", err)
        return
    }
    fmt.Printf("Output: %s\n",string(output))
    		/*
		⚡ Backend Analogy
Think of this as:
- Define request → exec.Command(...)
- Send request & get response → cmd.Output()
- Check for errors → if ERR != nil
- Print response body → fmt.Printf(...)
It’s like making an HTTP call to the OS shell and printing the response.

		*/
}