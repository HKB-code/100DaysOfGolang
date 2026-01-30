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


/*
Notes:
Q:when you execute a function as goRoutine its returned value is ignored
->Yes — when you start a function as a goroutine using the go keyword, any return value from that function is completely ignored.

Why the return value is ignored?
func calculateSum(a, b int) int {
    return a + b
}

func main() {
    go calculateSum(3, 4)          // ← return value (7) is thrown away
    // no way to get the 7 back directly!
}

When you write go f(...):
-A new goroutine is created
-The function f starts running concurrently
-Whatever f returns is discarded — there's no receiver for it
-Go does not provide any automatic way to capture the return value

How do we actually get values back from goroutines?
-There are two main idiomatic ways in Go:

1. Using channels (most common & recommended)

func calculateSum(a, b int, result chan<- int) {
    result <- a + b          // send result through channel
}

func main() {
    ch := make(chan int)
    
    go calculateSum(3, 4, ch)
    
    sum := <-ch              // receive the value
    fmt.Println(sum)         // 7
}
Or more cleanly (with return value):

func calculateSum(a, b int) int {
    return a + b
}

func main() {
    ch := make(chan int)
    
    go func() {
        ch <- calculateSum(3, 4)
    }()
    
    sum := <-ch
    fmt.Println(sum)
}


2. Using shared memory + synchronization (mutex, atomic, etc.)
(less preferred in modern Go)

var result int
var mu sync.Mutex

func calculateSum(a, b int) {
    mu.Lock()
    result = a + b
    mu.Unlock()
}

→ Channels are usually much cleaner and safer.


Most Important Takeaway (memorize this!)
-When you launch a function with go, it becomes fire-and-forget regarding its return value.
-If you care about the result → you must use a channel (or other synchronization mechanism).
-This is intentional design in Go — it forces you to think about communication between goroutines explicitly.



*/