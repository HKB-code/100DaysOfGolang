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
/*The Bigger Picture
The whole setup exists because:
- Go itself doesn’t know shell syntax — it only knows how to run executables.
- Many useful commands (echo, dir, copy) are not executables but shell built-ins.
- Wrapping them with "cmd /C" (Windows) or "bash -c" (Linux/macOS) gives Go access to the shell’s features.
- Capturing output (cmd.Output()) lets you use command results inside your Go program, instead of just printing them blindly.
So the purpose of “all these” is to bridge Go with the operating system’s shell, giving you a way to run commands, capture their output, and handle errors in a structured way.

/////////////////
The purpose of Go’s os/exec package is to give Go programs the ability to start and interact with external processes, including shell commands.
Here’s the bigger picture:
Why the exec package exists
- Run external programs: Go itself can’t directly run system commands; exec bridges that gap.
- Access shell features: Many commands (echo, dir, ls, cp) are shell built-ins, not standalone executables. exec lets you invoke them through the system shell.
- Capture output: Instead of just printing to the terminal, you can capture stdout/stderr and use it inside your Go logic.
- Control processes: You can start, stop, and manage external processes programmatically.
- Cross-platform abstraction: Same Go code can run commands on Windows, Linux, or macOS, with only minor differences in how you call the shell.
Example Purposes in Practice
- Automation: Run scripts or system utilities from Go.
- Integration: Call legacy tools or other languages’ compilers/interpreters.
- Data processing: Pipe command output into your Go program for further work.
- System tasks: Launch editors, open files, or manage OS-level operations.
Analogy
Think of exec as Go’s “remote control” for the operating system.
- Without it, Go is confined to its own runtime.
- With it, Go can say: “Hey OS, run this command for me, and give me back the result so I can continue working.”

/////////////////
Go cannot directly interpret shell commands like echo, dir, or ls because those aren’t standalone programs; they’re built-ins of the operating system’s shell. Go only knows how to run actual executables (like ping.exe, notepad.exe, or /usr/bin/curl).
The os/exec package was built to solve this gap:
- It lets Go launch external processes (executables or shell interpreters).
- By invoking the shell (cmd /C on Windows, bash -c on Linux/macOS), Go can indirectly run shell commands and built-ins.
- It provides structured ways to capture output, handle errors, and manage processes — things you can’t do by just “calling the shell” blindly.
So without exec, Go would be limited to its own runtime and couldn’t interact with the OS environment. With exec, Go becomes capable of automation, scripting, and system integration — essentially acting like a glue between Go code and the operating system.

👉 To make it concrete:
- Running exec.Command("ping", "google.com") works directly because ping is an executable.
- Running exec.Command("echo", "Hello") fails on Windows because echo isn’t an executable — it only exists inside cmd.exe. That’s why you need "cmd", "/C", "echo Hello".

		*/
}