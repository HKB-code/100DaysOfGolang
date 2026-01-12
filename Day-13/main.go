package main

import (
	"fmt"
	"sync"
	"time"
)

func Greet(phrase string){
	fmt.Println("Hello "+phrase)
}
func slowGreet(phrase string){
	time.Sleep(3*time.Second)
	fmt.Println(phrase)
}

func main() {
	var wg sync.WaitGroup

	wg.Add(4)
	go func(){
		defer wg.Done()
		Greet("World!")
	}()
go func(){
		defer wg.Done()
		Greet("My World!")
	}()
	go func(){
		defer wg.Done()
		slowGreet("How...are...you!")
	}()
	go func(){
		defer wg.Done()
		 Greet("opps")
	}()
	
	wg.Wait()
	
}