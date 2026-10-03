package main

import (
	"fmt"
	"time"
)

func PrintMessage(msg string) {
	for i := 1; i <= 3; i++ {
		fmt.Println(msg, ":", i)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go PrintMessage("Goroutine Background")

	PrintMessage("main thraed")
}
