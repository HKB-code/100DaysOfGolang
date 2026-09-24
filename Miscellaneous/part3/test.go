// If you really wanted to catch a panic
// Idiomatic Go: you avoid panics for normal control flow. But just to show you how it works:
package main

import "fmt"

func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("recovered from panic:", r)
        }
    }()

    a := 10
    b := 0
    fmt.Println(a / b) // this will panic, but we recover
}
/*
Step-by-step execution of your program
- Program enters main.
- defer func() { ... }() is registered to run when main ends (or panics).
- a := 10, b := 0.
- fmt.Println(a / b) executes:
- Division by zero → panic (runtime error: integer divide by zero).
- Panic starts unwinding:
- Before main fully exits, Go runs deferred functions.
- Deferred function runs:
- Calls recover(), which catches the panic.
- r now holds something like "runtime error: integer divide by zero".
- Prints: recovered from panic: runtime error: integer divide by zero.
- Since the panic is recovered, the program does not crash. main completes gracefully after the deferred function.
No stack trace, no fatal crash — just your custom output.
//////////////////////////
When should you use this pattern?
Important nuance: this is a tool, not a pattern for normal control flow.
Good uses:
- Top-level protection in main:
- So unexpected panics don’t crash the whole process without logging something nice.
- Server handlers / goroutines:
- Recover inside a goroutine so one panic doesn’t kill the entire server.
- Libraries:
- If you might panic internally but want to expose errors to callers instead.
Bad uses:
- As a substitute for if err != nil style error handling.
- For normal conditions like divide by zero that you could just check with if b == 0.
Think of panic/recover like a fire extinguisher, not your daily cooking tool
////////////////////////////
A super clear mental model
- panic: “Abort! Something is wrong.”
- defer: “Run this later, even if things go wrong.”
- recover: “Catch and neutralize the panic so the program doesn’t die.”
Your code is essentially saying:
“Try to run this code. If it panics for any reason (like division by zero), don’t crash; just log that we recovered from a panic.”




*/