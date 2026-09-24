//real‑world Go example: an HTTP server where a single panicking handler does NOT crash the entire server. This is exactly how production Go servers protect themselves.

/*
✅ Real‑World Example: Safe HTTP Handler With Panic Recovery
🔥 Problem
If any handler panics, the entire server would normally crash.
✅ Solution
Wrap each handler with a middleware that uses defer + recover to catch panics and return a safe response.

*/

package main

import (
	"fmt"
	"log"
	"net/http"
)

// Middleware that recovers from panics inside handlers
func safeHandler(h http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if err := recover(); err != nil {
                log.Println("Recovered from panic:", err)
                http.Error(w, "Internal Server Error", http.StatusInternalServerError)
            }
        }()
        h(w, r)
    }
}

func main() {
    // This handler intentionally panics
    http.HandleFunc("/panic", safeHandler(func(w http.ResponseWriter, r *http.Request) {
        panic("something went wrong in handler")
    }))

    // Normal handler
    http.HandleFunc("/", safeHandler(func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Hello, world!")
    }))

    fmt.Println("Server running on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

/*
✅ What Happens Here
✅ 1. safeHandler wraps every handler
It adds a defer block that catches panics:

defer func() {
    if err := recover(); err != nil {
        log.Println("Recovered:", err)
        http.Error(w, "Internal Server Error", 500)
    }
}()

✅ 2. If a handler panics
Example:
panic("something went wrong")

Then:
- The panic is caught
- The server logs the error
- The client gets a 500 Internal Server Error
- The server keeps running
✅ 3. Other requests continue normally
One bad request does not kill the server.

✅ Why This Matters in Real Servers
This pattern is used in:
- Production Go APIs
- Microservices
- Web frameworks
- Background workers
- Goroutines
It ensures:
- Fault isolation
- No global crashes
- Graceful error handling

*/