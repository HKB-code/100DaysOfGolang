package main

import (
	"fmt"
	"time"
)


func Greet(phrase string,doneChan chan bool){
	fmt.Println("Hello "+phrase)
	doneChan<-true
}
func slowGreet(phrase string,doneChan chan bool){
	time.Sleep(3*time.Second)
	fmt.Println(phrase)
	doneChan<-true
	close(doneChan)
}

func main() {
	// dones:= make([]chan bool,4)
	// dones[0] =  make(chan bool)
	done:= make(chan bool)
 go	Greet("world",done)
//  dones[1] =  make(chan bool)
	go Greet("How are you?",done)
//    dones[2] =  make(chan bool)
	go slowGreet("my friend",done)
	// dones[3] =  make(chan bool)
	go Greet("Jhon",done)
// for _,done:= range dones{
// 	<-done
// }
// for doneChan:= range done{
// 	fmt.Println(doneChan)
// }
// }
for range done{

}
}