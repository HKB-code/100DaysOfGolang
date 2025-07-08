/*
we'll have to add this file to a package

because every Go file must be part of a package.

# That's how the different files are then linked together

behind the scenes by the Go compiler.
*/
package main

import "fmt"

func options(){
	fmt.Println("What do you want to do?")
	fmt.Println("1.Check Balance")
	fmt.Println("2.Deposit Money")
	fmt.Println("3.Withdraw Money")
	fmt.Println("4.Exit")
}

/*

You need to run both files, else they are not automatically associated with each other during runtime (this is why we often use submodules as imports instead, since they do "magically" associate with the running program file).

btw, a shortcut to run all files in a folder is go run .
*/