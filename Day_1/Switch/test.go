package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("Go runs on")

	switch os:= runtime.GOOS;os{
	case "darwin":
		fmt.Println("macOS.")
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Printf("%s\n",os )
	}
	today:= time.Now().Weekday()
	fmt.Println("When's Saturday?")
	switch time.Saturday{
	case today+0:
		fmt.Println("today")
	case today+1:
		fmt.Println("tomorrow")
	case today + 2:
		fmt.Println("In two days.")
	
	default:
		fmt.Println("Too far away")
	}
	
}