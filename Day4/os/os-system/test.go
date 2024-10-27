package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

func main() {
	// 6. Get System Information
    fmt.Println("Go Version:", runtime.Version())
    fmt.Println("OS:", runtime.GOOS)
    fmt.Println("Architecture:", runtime.GOARCH)


// 7. Execute a Shell Command
		cmd := exec.Command("echo", "Hello, World!")
    output, err := cmd.Output()
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Printf("Output: %s\n", output)
}
