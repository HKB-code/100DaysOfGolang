/*
The flag package in Go (Golang) is part of the standard library and implements command-line flag parsing. It allows Go programs to define, parse, and access optional command-line arguments (flags) that customize the program's behavior, such as configuration options passed when running the executable.

What Are Command-Line Flags?
Command-line flags (also called options or switches) are arguments passed to a program when it's run from the terminal. They typically start with - or -- and allow users to customize the program's behavior without changing the code.

Examples:
-ls -l (the -l is a flag for long listing)
-go run main.go -name=Grok (custom flag)

In Go, the built-in flag package makes it easy to define, parse, and use these flags in your programs.

Key Features
-Supports built-in types: strings, integers, booleans, floats, durations, etc.
-Provides default values and usage descriptions for each flag.
-Automatically generates help/usage messages (e.g., via -h or -help).
-Allows access to remaining non-flag arguments after parsing.
-Supports custom flag types by implementing the Value interface.
-Enables subcommands or multiple flag sets via FlagSet.

Basic Usage
-Define flags (often at the package level).
-Call flag.Parse() in main() to process os.Args.
-Use the flag values (they are pointers) or bind them directly to variables.

Advanced FeaturesCustom Types:
-Implement the flag.Value interface (methods: String() and Set(string) error).
-Subcommands: Use flag.NewFlagSet() for tools like git commit -m "msg".
-Positional Arguments: Access with flag.Args() (slice) or flag.Arg(i) (single).
-Print Defaults: flag.PrintDefaults() for custom help.
*/
package main

import (
	"flag"
	"fmt"
	"strings"
)

// It satisfy the flag.Value Interface,That's the beauty of Go interfaces. No registration, no boilerplate – just implement the methods and it "just works".

type Strings []string
 /*
 once you correctly define the two required methods on the pointer receiver (*Strings),
Go and the flag package automatically use them whenever needed.


 */
 func (s *Strings)String()string{return  strings.Join(*s,",")}
 /*
 func (s *Strings) Set(v string) error {
    // Agar comma hai to split karo, warna poora as one item add karo
    if strings.Contains(v, ",") {
        parts := strings.Split(v, ",")
        for _, part := range parts {
            trimmed := strings.TrimSpace(part)
            if trimmed != "" {
                *s = append(*s, trimmed)
            }
        }
    } else {
        *s = append(*s, v)
    }
    return nil
}
 */
//  best practice
 func(s *Strings)Set(v string)error{
	*s = append(*s,v)
	return  nil
 }
func main() {
	// define a flag
	name:= flag.String("name","world","a name to greet")
	count := flag.Int("count", 1, "How many times to greet")
    verbose := flag.Bool("verbose", false, "Enable verbose output")
	var quantity int
	cream := flag.Bool("Cream",false,"add cream?")

	// Bind directly to a variable 
	flag.IntVar(&quantity, "quantity", 2,"number of shots")

	  // Define the custom repeatable flag
    var tags Strings
    flag.Var(&tags, "tag", "add a tag (can be repeated)")

// parse command-line arguments
	flag.Parse()
	// use flag value
	
	for i:=0;i<*count;i++{
fmt.Printf("Hello, %s!\n",*name)
if *verbose{
	 fmt.Println("Greeting number:", i+1)
}
	}

	    if len(tags) == 0 {
        fmt.Println("No tags provided")
    } else {
        fmt.Println("Tags:")
        for i, tag := range tags {
            fmt.Printf("  %d. %s\n", i+1, tag)
        }
    }

	  // Positional arguments (anything after flags)
	//   "NArg is the number of arguments remaining after flags have been processed."


    if flag.NArg() > 0 {
        fmt.Println("Extra arguments:", flag.Args())
    }


	fmt.Println("Quantity:", quantity)
    fmt.Println("Cream:", *cream)

		
}