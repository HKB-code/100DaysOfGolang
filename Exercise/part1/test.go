package main

import "fmt"

func main(){
	adam:=42
	fmt.Printf("42 as binary, %b \n",adam)
	fmt.Printf("42 as hexadecimal, %x \n",adam)
	fmt.Printf("42 as hexadecimal, %X \n",adam)
fmt.Println("------------------")
fmt.Printf("42 as binary, %#b \n",adam)
	fmt.Printf("42 as hexadecimal, %#x \n",adam)



	// print these values as both binary and hexadecimal
	a,b,c,d,e,f:= 0,1,2,3,4,5

	fmt.Printf("%v\t%b\t%X\n", a,a,a)
	fmt.Printf("%v\t%b\t%X\n", b,b,b)
	fmt.Printf("%v\t%b\t%X\n", c,c,c)
	fmt.Printf("%v\t%b\t%X\n", d,d,d)
	fmt.Printf("%v\t%b\t%X\n", e,e,e)
	fmt.Printf("%v\t%b\t%X\n", f,f,f)


}

/*
🔑 Difference between %b and %#b
- %b
Prints the value in binary format (base 2).

- %#b
Prints the value in binary format with a base prefix.
- In Go, the prefix for binary is 0b.

Note:
-So the # flag adds the base indicator (0b for binary, 0x for hex, 0 for octal)



-Base 2 is simply another way of saying binary — the number system that uses only two digits: 0 and 1.

🧮 Key Points
- Base means the number of unique digits a system uses.
- Base 10 (decimal) → digits 0–9
- Base 2 (binary) → digits 0–1
- Each position in a binary number represents a power of 2, just like in decimal each position represents a power of 10.

🔢 Example: Number 42 in Base 2
- Decimal (base 10): 42
- Binary (base 2): 101010
Breakdown:
42=1\cdot 2^5+0\cdot 2^4+1\cdot 2^3+0\cdot 2^2+1\cdot 2^1+0\cdot 2^0
42=32+8+2

-⚡ Why It Matters
- Computers use base 2 internally because electronic circuits have two states: on (1) and off (0).
- All data — text, images, programs — is ultimately stored and processed as binary.




*/