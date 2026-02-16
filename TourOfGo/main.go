// package main

// import "fmt"

// func main(){
// 	fmt.Println("Hello")
// 	a:=1
// 	b:=2
// 	res:=add(a,b)
// 	fmt.Println(res)

// for i:=1;i<10;i++{
// 	fmt.Println(i)
// }

// num:=[]int{1,2,3,4,5}
// for i,v:=range num{
// 	fmt.Printf("Index: %d, Value: %d\n", i,v)
// }

// k:=map[string]int{"a":1,"b":2,"c":3}

// for key,value:=range k{
// 	fmt.Printf("Key: %s, Value: %d\n", key,value)
// }

// w:=1
// for w<4{
// 	fmt.Println(w)
// 	w++
// }

// }

// func add(a,b int)int{
// 	return  a+b
// }
/////////////////////////////////////////////////

// package main

// import (
// 	"fmt"
// 	"math"
// )

// const inflamationRate = 2.5

// func main(){
// 	investmentAmount:= getText("Enter investment amount: ")
// 	expectedAmount:=getText("Expected Amount: ")
// 	years:= getText("Years")

// futureValue,realFutureValue:= calculateFutureValue(investmentAmount,expectedAmount,years)

// fmt.Println(futureValue)
// fmt.Println(realFutureValue)

// fmt.Printf("Future Value %.2f\n",futureValue)
// fmt.Printf("Real Future Value %.2f\n",realFutureValue)
// }

// func getText(prompt string)float64{
// 	var value float64
// 	fmt.Println(prompt)
// 	fmt.Scan(&value)
// 	return  value
// }

// func calculateFutureValue(investmentAmount,expectedAmount,years float64)(fv,rfv float64){
// 	fv = investmentAmount*math.Pow(1+expectedAmount/100,years)
// 	rfv = fv/math.Pow(1+inflamationRate/100,years)
// 	return
// }

///////////////////////////////////////////////

// package main

// import "fmt"

// func main(){
// revenue:=getText("Revenue: ")
// expenses:= getText("Expenses: ")
// taxRate:= getText("Tax Rate: ")

// ebt,profit,ratio := calc(revenue,expenses,taxRate)
// fmt.Printf("ebt: %.2f\n",ebt)
// fmt.Printf("Profit: %.2f\n",profit)
// fmt.Printf("Ratio: %.2f\n",ratio)

// }

// func getText(prompt string) float64{
// 	var value float64
// 	fmt.Println(prompt)
// 	fmt.Scan(&value)
// 	return  value
// }

// func calc(revenue,expenses,taxRate float64)(ebt,profit,ratio float64){
// 	ebt= revenue-expenses
// 	profit = ebt*(1-taxRate/100)
// 	ratio = ebt/profit
// return
// }

// ////////////////////////
// package main

// import (
// 	"fmt"
// 	"math"
// )

// func pow(x,n,lim float64)float64{
// 	if v:=math.Pow(x,n);v<lim{
// 		return v
// 	}
// 	return  lim
// }

// func main(){
// 	fmt.Println(pow(3,2,10),pow(3,3,20))
// }

// /////////////////////////////////////
// package main

// import (
// 	"fmt"
// 	"runtime"
// 	"time"
// )

// func main(){
// 	switch os:=runtime.GOOS; os{
// 	case "darwin":
// 	 fmt.Println("macOs")

// 	case "linux":
// 		fmt.Println("Linux")
// 	default:
// 		fmt.Printf("%s\n",os)

// 	}

// 	today := time.Now().Weekday()
// 	fmt.Println(today,time.Monday)

// 	switch today{
// 	case time.Saturday, time.Sunday:
// 		fmt.Println("Weekend")
// 	default:
// 		fmt.Println("Weekday: ", today)

// 	}

// switch time.Saturday{
// case today+0:
// 	fmt.Println("Today")
// case today+1:
// 	fmt.Println("Tomorrow")
// case today+2:
// 	fmt.Println("In two days.")
// default:
// 		fmt.Println("Too far away.")
// }

// /*
// Switch without a condition is the same as switch true.

// This construct can be a clean way to write long if-then-else chains.
// */

// t:=time.Now()
// switch{
// case t.Hour()<12:
// 	fmt.Println("GM")
// case t.Hour()<17:
// 	fmt.Println("GA")
// default:
// 	fmt.Println("Good evening.")
// }
// /*
// Defer
// A defer statement defers the execution of a function until the surrounding function returns.

// The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

// Stacking defers
// Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.
// */
// fmt.Println("counting")

// 	for i := 0; i < 10; i++ {
// 		defer fmt.Println(i)
// 	}

// 	fmt.Println("done")

// }

///////////

// package main

// import "fmt"

// func main(){

// 	i,j :=22,43
// fmt.Println(j)
// 	p:=&i
// 	fmt.Println(p) // read memory address
// 	fmt.Println(*p) //read the value

// 	// dereferncing
// 	*p = 55
// 	fmt.Println(i)

// 	p = &j
// 	fmt.Println(p)
// 	*p = *p/2
// 	fmt.Println(j)

// 	//////

// }
///////////////////////
// struct: A struct is a collection of fields.

// package main

// import "fmt"

// type Vertex struct{
// 	x int
// 	y int
// }
// var(
// 	v1 = Vertex{1,2}
// 	v2 =Vertex{x: 7}
// 	v3 =Vertex{}

// 	z = &v1
// )

// func main(){
// 	v := Vertex{1,2}
// 	fmt.Println(v.x)
// 	p:=&v
// 	// (*p).x = 12
// 	p.x = 12
// fmt.Println(v)

// fmt.Println(v1,v2,v3,z)
// fmt.Println(v1,v2,v3,*z)

// /*
// An array's length is part of its type, so arrays cannot be resized.
// */
// var a[2]int
// a[0] =1
// a[1] = 2
// fmt.Println(a)

// /*
// An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array. In practice, slices are much more common than arrays.

// The type []T is a slice with elements of type T.
// */
// primes:=[5]int{2,3,5,7,11}

// var s []int= primes[1:4]
// fmt.Println(s)

// /*

// Slices are like references to arrays
// A slice does not store any data, it just describes a section of an underlying array.

// Changing the elements of a slice modifies the corresponding elements of its underlying array.

// Other slices that share the same underlying array will see those changes.
// */

// names:= [3]string{
// 	"aa",
// 	"bb",
// 	"cc"}

// k:=names[0:2]
// fmt.Println(k)
// k[1] = "dd"
// fmt.Println(names)

// a1:= []struct{
// 	i int
// 	b bool

// }{
// 	{1, false},
// 	{2,true},
// 	{3,false},
// 	{4,true}}

// fmt.Println(a1)

// a1= a1[0:2]
// fmt.Println(a1)

// a1 = a1[1:2]
// fmt.Println(a1)

// l:=[]int{
// 	1,2,3,4,5,6}

// fmt.Println(l)
// printSlices(l)

// l = l[0:3]
// fmt.Println(l)
// printSlices(l)

// l=l[2:3]
// fmt.Println(l)
// printSlices(l)

// //Nil slices

// /*
// The zero value of a slice is nil.

// A nil slice has a length and capacity of 0 and has no underlying array.
// */

// var x []int
// fmt.Println(x,len(x),cap(x))
// if x==nil{
// 	fmt.Println("NIL!")
// }
// }

// func printSlices(m []int){
// 	/*
// 	A slice has both a length and a capacity.

// The length of a slice is the number of elements it contains.

// The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.

// The length and capacity of a slice s can be obtained using the expressions len(s) and cap(s).

// You can extend a slice's length by re-slicing it, provided it has sufficient capacity. Try changing one of the slice operations in the example program to extend it beyond its capacity and see what happens.
// 	*/
// 	fmt.Printf("len=%d cap=%d %v\n", len(m), cap(m), m)
// }

// ////////////////////////////////////////////////////////
/*

Creating a slice with make
Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.

The make function allocates a zeroed array and returns a slice that refers to that array:
*/
// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main(){
// 	a:= make([]int,0,5)
// 	fmt.Println(a, len(a),cap(a))
// 	a=a[:cap(a)]
// 	fmt.Println(a, len(a))
// 	a=a[1:]
// 	fmt.Println(a, len(a))

// 	/*

// Slices of slices
// Slices can contain any type, including other slices.
// 	*/

// board := [][]string{
//         []string{"_", "_", "_"},
// 		[]string{"_", "_", "_"},
// 		[]string{"_", "_", "_"},
// }
// fmt.Println(board)
// // player take turn
// board[0][0] = "X"
// board[0][2] = "x"
// board[1][0]="O"
// board[1][2] = "x"
// board[2][2]="O"

// for i:=0;i<len(board);i++{
// 	fmt.Printf("%s\n",strings.Join(board[i]," "))
// }

// /////
// /*
// 🔎 Slice Growth Rules in Go
// - A slice in Go has length (number of elements) and capacity (size of the underlying array).
// - When you append, if the slice has enough capacity, it just extends the length.
// - If not, Go allocates a new underlying array with a larger capacity (usually doubling, but the exact growth strategy can vary depending on implementation and size).

// */
// k:= []int{}
// printSlice(k)

// k = append(k, 0)
// printSlice(k)

// k = append(k, 1)
// printSlice(k)

// k = append(k, 2,3,4,5,6)
// printSlice(k)

// var pow = make([]int,10)
// for i := range pow{
// 	pow[i] = 1<<uint(i)
// }

// for _,value:= range pow{
// 	fmt.Printf("%d\n",value)
// }

// }
// func printSlice(s []int) {
// 	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
// }

// ////////////////////////////////
// // maps
// package main

// import "fmt"

// type Vertex struct{
// 	Lat,Long float64
// }

// var m map[string]Vertex

// var z  = map[string]Vertex{
// 	"a": Vertex{
// 		12,1334,
// 	},
// 	"b":Vertex{
// 		456,345667,
// 	},
// }

// /*
// Map literals continued
// If the top-level type is just a type name, you can omit it from the elements of the literal.
// */
// var x = map[string]Vertex{
// 	"a": {
// 		987,5454,
// 	},
// 	"b": {
// 		5645,8675,
// 	},
// }

// func main() {
// 	m = make(map[string]Vertex)
// 	m["Bell Labs"] = Vertex{
// 		40.68433, -74.39967,
// 	}
// 	fmt.Println(m["Bell Labs"])

// 	/*
// 	Map literals
// Map literals are like struct literals, but the keys are required.
// 	*/
// fmt.Println(z)
// fmt.Println(x)

// // Mutating Maps

// d:= make(map[string]int)

// d["answer"] = 42
// 	fmt.Println("The value:", d["answer"])

// d["answer"] = 48
// 	fmt.Println("The value:", d["answer"])

// 	delete(d, "answer")
// 	fmt.Println("The value:", d["answer"])

// h,ok:=d["answer"]
// fmt.Println("The value:", h, "Present?", ok)
// }

//////////////////////////////
// Function Values

// package main

// import (
// 	"fmt"
// 	"math"
// )

// /*
// Functions are values too. They can be passed around just like other values.

// Function values may be used as function arguments and return values.
// */

// func compute(fn func(float64 ,float64)float64)float64{
// 	return  fn(3,4)
// }

// func main(){
// 	hypot:= func (x float64,y float64)float64  {
// 		return math.Sqrt(x*x+y*y)
// 	}

// 	fmt.Println(hypot(5,12))

// 	fmt.Println(compute(hypot))

// 	fmt.Println(compute(math.Pow))
// }

/////////////////////////
/*
A closure in Go is simply a function that “remembers” variables from the scope in which it was created, even after that scope has finished executing.

🧩 What is a Closure?
- In Go, you can define anonymous functions (functions without names).
- When such a function captures variables from its surrounding scope, it becomes a closure.
- This means the function can access and modify those variables later, even if the outer function has already returned

⚡ Key Properties of Closures
- Capture environment: They hold onto variables from the outer scope.
- Stateful functions: They can maintain state across multiple calls.
- Encapsulation: Useful for hiding implementation details while exposing controlled behavior.
- Independent instances: Each closure has its own captured environment.

🚀 Practical Use Cases
- Counters: As shown above.
- Factories: Functions that return customized functions.
- Event handlers: Retain context/state for callbacks.
- Encapsulation: Hide internal state without using global variables.

In short: A closure is a function bundled with its environment. It’s Go’s way of letting functions carry state around elegantly.



*/

// package main

// import "fmt"

// func adder() func(int)int{

// sum:=0
// 	return func (i int) int{
// sum+=i
// return  sum
// 	}
// }

// func main(){
// pos,neg := adder(),adder()

// for i:=0;i<10;i++{
// fmt.Println(pos(i),neg(-2*i))
// }
// }

//////////////////////////////
// // Methods:
// /*
// 🔑 Key Points About Methods in Go
// - No classes: Go doesn’t have classes like Java or C++.
// - Receiver: Instead, methods are defined on types using a receiver argument.
// - Syntax: The receiver is placed between the func keyword and the method name.
// - Value vs Pointer Receiver:
// - A value receiver works on a copy of the type.
// - A pointer receiver allows modifying the original value.

// */
// package main

// import (
// 	"fmt"
// 	"math"
// )

// type Vertex struct {
// 	X,Y float64
// }
// /*
// 📌 Rule Recap
// - You can only declare a method on a type you define in the same package.
// - You cannot attach methods to:
// - Built‑in types (int, string, etc.)
// - Types defined in another package (like time.Time from the time package).

// 🧩 Why This Restriction Exists
// - Encapsulation & Ownership: Methods are meant to extend the behavior of types you own. If Go allowed attaching methods to foreign types, it could cause conflicts (two packages defining different methods on the same type).
// - Simplicity: Keeps Go’s type system clean and avoids “monkey patching” chaos you sometimes see in dynamic languages.

// */
// type myFloat float64

// func (v Vertex) Abs()float64{
// 	return  math.Sqrt(v.X*v.X+v.Y*v.Y)
// }

// func (f myFloat)s()float64{
// 	if f<0{
// 		return float64(-f)
// 	}
// 	return float64(f)
// }
// /////////////
// /*
// Methods with pointer receivers can modify the value to which the receiver points (as Scale does here). Since methods often need to modify their receiver, pointer receivers are more common than value receivers.
// */

// type voides struct{
// 	X,Y float64
// }

// func (k voides) Sqrt()float64{
// 	return  math.Sqrt(k.X*k.X+k.Y*k.Y)
// }
// /*
// ⚖️ When to Use Which
// - Value receiver:
// - Small, immutable types.
// - Methods that don’t modify the receiver.
// - Example: Abs().

// - Pointer receiver:
// - Large structs (avoid copying overhead).
// - Methods that need to mutate the receiver.
// - Example: Scale().

// */

// func (k *voides)scale(f float64){
// k.X = k.X*f
// k.Y = k.Y*f
// }

// func main(){
// v:= Vertex{
// 4,3}
// fmt.Println(v.Abs())

// f:= myFloat(-math.Sqrt(16))

// fmt.Println(f.s())
// var z myFloat = 2.0
// fmt.Println(z)

// //pointer
// K:= voides{3,4}
// // (&K).scale(10)
// /*
//  the method with the pointer receiver is called automatically. That is, as a convenience, Go interprets the statement v.Scale(5) as (&v).Scale(5) since the Scale method has a pointer receiver.
// */
// K.scale(10)
// fmt.Println(K.Sqrt())

// }
////////////////////////
// package main

// import "fmt"

// //  Methods and pointer indirection(1)

// type Vertex struct {
// 	x,y float64
// }

// func (v *Vertex) scale(f float64){
// 	v.x *=f
// 	v.y *=f
// }

// func scaleFunc(v *Vertex, f float64){
// 		v.x *=f
// 	v.y *=f

// }

// func main(){

// 	v:=Vertex{3,4}
//    v.scale(10)
//    scaleFunc(&v,10)
//    fmt.Println(v)

//    p:= &v
//    p.scale(20)
//    scaleFunc(p,20)
//       fmt.Println(p)
// }
///////////

//  Methods and pointer indirection(2)

// package main

// import (
// 	"fmt"
// 	"math"
// )

// type Vertex struct {
// 	X, Y float64
// }

// func (v Vertex) Abs() float64 {
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }

// func AbsFunc(v Vertex) float64 {
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }

// func main() {
// 	v := Vertex{3, 4}
// 	fmt.Println(v.Abs())
// 	fmt.Println(AbsFunc(v))

// p := &Vertex{4, 3}
// 	fmt.Println(p.Abs())
// 	fmt.Println(AbsFunc(*p))

// }
/*
🔹 Why the difference?
- Methods: Go adds syntactic sugar. If the receiver type is Vertex, you can call it on both Vertex and *Vertex — Go will auto-dereference or auto-address as needed.
- Functions: No sugar. You must match the parameter type exactly. If it wants Vertex, you give it a Vertex. If you have *Vertex, you dereference manually.

✅ Summary:
- Methods: Go automatically handles pointer ↔ value conversions when calling them.
- Functions: You must explicitly dereference (*p) or take the address (&v) to match the parameter type.

🖥️Choosing a value or pointer receiver
There are two reasons to use a pointer receiver.

The first is so that the method can modify the value that its receiver points to.

The second is to avoid copying the value on each method call. This can be more efficient if the receiver is a large struct, for example.
*/

////////////////////////////

// Interface
//part-1

// package main

// import (
// 	"fmt"
// 	"math"
// )

// type Abser interface{
// 	Abs() float64
// }

// type myFloat float64

// func (f myFloat) Abs ()float64{
// 	if f<0{
// 		return  float64(-f)
// 	}
// 	return  float64(f)
// }

// type Vertex struct{x,y float64}

// func (v *Vertex) Abs() float64{
// 	return  math.Sqrt(v.x*v.x+v.y*v.y)
// }

// func main(){
// 	var a Abser
// 	fmt.Println(a)
// 	f:=myFloat(-math.Sqrt(2))
// 	v:=Vertex{3,4}

// 	a=f
// 	fmt.Println(a.Abs())
// 	a= &v
// 	fmt.Println(a.Abs())
// }
///////////////
//part-2

// package main

// import "fmt"

// type I interface {
// 	M()
// }

// type T struct{
// 	S string
// }

// // This method means type T implements the interface I,
// // but we don't need to explicitly declare that it does so.

// func (t T) M(){
// 	fmt.Println(t.S)
// }

// func main(){
// 	var i I = T{
// 		"Hello",
// 	}
// i.M()

// }

//
//part-3
/*
🔎 What’s Inside an Interface Value
When you assign a concrete type to an interface variable, Go stores it as a pair:
(\mathrm{value},\mathrm{type})
- Value → the actual data (e.g., a Dog{} struct instance).
- Type → the concrete type information (e.g., Dog).
So an interface variable isn’t just holding the value — it also remembers what type that value is.

*/
// package main

// import (
// 	"fmt"
// 	"math"
// )

// type I interface{
// 	M()
// }
// func describe(i I) {
// 	fmt.Printf("(%v, %T)\n", i, i)
// }

// type T struct {
// 	S string
// }

// func (t *T) M(){
// 	fmt.Println(t.S)
// }

// type F float64

// func (f F) M() {
// 	fmt.Println(f)
// }

// func main(){
// 	var i I

// 	t:= T{"Hello"}
// 	t.M()
// 	i = &t
// 	describe(i)
// 	i.M()

// 	i = F(math.Pi)
// 	describe(i)
// 	i.M()
// }

// ////////////////
//  Interface values with nil underlying values

// package main

// /*
// Note that an interface value that holds a nil concrete value is itself non-nil.
// */
// import "fmt"

// type I interface{
// 	M()
// }

// type T struct{
// 	S string
// }

// func (t*T) M(){
// 	if t == nil{
// 		fmt.Println("<nil>")
// 		return
// 	}
// 	fmt.Println(t.S)
// }

// func describe(i I) {
// 	fmt.Printf("(%v, %T)\n", i, i)
// }

// func main(){
//  var i I

//  var t *T
//  i = t
//  describe(i)
//  i.M()

//  	i = &T{"hello"}
// 	describe(i)
// 	i.M()
// 	/*
// 	🔎 Interface value with nil concrete value
// - Ek interface variable internally (value, type) store karta hai.
// - Agar tum ek nil concrete value assign karte ho (jaise var t *T = nil aur phir i = t), to interface ke andar (value=nil, type=*T) store hota hai.
// - Is case me interface khud nil nahi hota, kyunki uske paas type information abhi bhi hai (*T).
// 👉 Matlab: interface ke andar type set hai, bas us type ka value nil hai

// 	*/
// }
///////////////////////////

//Nil interface values
/*A nil interface value holds neither value nor concrete type.

Calling a method on a nil interface is a run-time error because there is no type inside the interface tuple to indicate which concrete method to call.*/

// package main

// import "fmt"

// type I interface {
// 	M()
// }

// func describe(i I) {
// 	fmt.Printf("(%v, %T)\n", i, i)
// }

// func main(){
// 	var i I
// 	describe(i)
// 	i.M()
// }
////////////////////////
/*

The empty interface
The interface type that specifies zero methods is known as the empty interface:

interface{}
An empty interface may hold values of any type. (Every type implements at least zero methods.)

Empty interfaces are used by code that handles values of unknown type. For example, fmt.Print takes any number of arguments of type interface{}.
*/

// package main

// import "fmt"

// func describe(i interface{}) {
// 	fmt.Printf("(%v, %T)\n", i, i)
// }

// func main(){
// 	var i interface {}
// 	describe(i)
// 	var z any
// 	describe(z)

// 	i =42
// 	z=43
// 	describe(i)
// 	describe(z)

// 	i ="Hello"
// 	z="Hey"
// 	describe(i)
// 	describe(z)
// }

/////////
// Type assertions
/*
🔑 Key Points
- Concrete type → The actual type of value stored inside an interface (like int, string, etc.).
- Type assertion → A way to take the value out of an interface and use it as its real type.
- Syntax:
- t := i.(T) → Unsafe form. If i is not type T, program panics.
- t, ok := i.(T) → Safe form.
- If i is type T → ok = true, t has the value.
- If not → ok = false, t is zero value of T.
- Analogy → Just like reading from a map:
- value, ok := m[key] → ok tells if the key exists.
- Same with type assertion → ok tells if the type matches.

👉 In short:
Interface = box.
Concrete type = actual item inside the box.
Type assertion = checking and pulling that item out safely.

*/

// package main

// import "fmt"

// func main(){
// 	// var i interface{}
// 	// fmt.Println(i)

// 		var i interface{} = "Hello"
// 	// fmt.Println(i)

// 	s:= i.(string)
// 	fmt.Println(s)

// 	s,ok:=i.(string)
// 	fmt.Println(s,ok)

// 	//safe form
// 		f, ok := i.(float64)
// 	fmt.Println(f, ok)

//unsafe form
// f = i.(float64) // panic
// fmt.Println(f)
/*
	🔹 Analogy
Think of it like opening a box:
- Safe form (f, ok) → You peek inside carefully. If the item isn’t what you expected, you just say “nope” and move on.
- Unsafe form (f = i.(T)) → You grab it blindly. If it’s the wrong item, you hurt yourself (panic).

👉 So the rule of thumb:
- Use f, ok := i.(T) when you’re not 100% sure of the type.
- Use f = i.(T) only when you’re absolutely certain the interface holds that type.

*/
// }

///////////////////////
/*
🔑 Type Switch (Quick Notes)
- Definition → A type switch lets you check multiple possible types of an interface value in one construct.

//example:
switch v := i.(type) {
case T:
    // v is of type T
case S:
    // v is of type S
default:
    // v has same type as i (interface)
}


- Key points:
- Works like a normal switch, but cases are types instead of values.
- i.(type) is used instead of i.(T) (keyword type replaces specific type).
- In each case, v is automatically converted to that concrete type.
- default runs if no type matches → v stays as the interface type.




👉 In short:
Type switch = cleaner way to handle multiple type assertions at once.

*/
// package main

// import "fmt"

// func do(i interface{}){
// 	switch v := i.(type){
// 	case int:
// 		fmt.Printf("Twice %v is %v\n", v,v*2)
// 	case string:
// 		fmt.Printf("%q is %v bytes long\n", v, len(v))
// 	default:
// 		fmt.Printf("I don't know about type %T!\n", v)

// 	}
// }
// func main() {
// 	do(21)
// 	do("hello")
// 	do(true)
// }

/////////////////////////
// stringers
/*
- Any type that implements String() string method is a Stringer.
- Purpose:
- Allows a type to describe itself as a string.
- Used by fmt package (like fmt.Println, fmt.Printf) to decide how to print values.
- Behavior:
- If a type implements String() string, fmt will call that method automatically when printing.
- If not, fmt falls back to default formatting.

👉 In short: Stringer = “self-description as string” interface.

- Key Point:
- Stringer makes custom types human-readable when printed.
- It’s one of the most common interfaces in Go.



*/
// package main

// import "fmt"

// type Person struct{
// 	Name string
// 	Age int
// }

// func (P Person) String() string{
// 	return  fmt.Sprintf("%v (%v years) ", P.Name,P.Age)
// }

// func main(){
// 	a := Person{"Harshi Jain",44}
// 	z := Person{"Zaphod Beeblebrox", 9001}
// 	fmt.Println(a, z)
// }

///////////
/*
🔑 Errors in Go
- Error type → Built-in interface:
type error interface {
    Error() string
}
- Any type that implements Error() string satisfies the error interface.
- Printing errors → Just like fmt.Stringer, the fmt package looks for Error() when printing values.
- Returning errors → Functions often return (value, error) pairs.
- Best practice → Always check err before using the returned value.

👉 In short:
Errors in Go are values.
They implement Error() string, and you handle them by checking if err != nil.

*/

// package main

// import (
// 	"fmt"
// 	"time"
// )

// type MyError struct{
// 	When time.Time
// 	What string
// }

// func (e* MyError) Error()string{
// 		return fmt.Sprintf("at %v, %s",
// 		e.When, e.What)
// }
/*
- In Go, error is an interface:
type error interface {
    Error() string
}
- For run() to return error, the value you return must implement this interface.
- That means: whatever you return must have an Error() string method.

*/

// func run()error{
// 	return &MyError{
// 		time.Now(),
// 		"it didn't work",
// 	}
// }

// func main(){
// 	if err:= run();err!=nil{
// 		fmt.Println(err)
// 	}
// }

// //////////////////
// Readers
/*
type Reader interface {
    Read(p []byte) (n int, err error)
}

Notes
- io.Reader is the read end of a stream abstraction.
- Any type that implements Read with this signature can be treated as a Reader.
- Read fills the slice p with data, returns how many bytes were read (n), and an error (err).
- When the stream ends, Read returns io.EOF.
Key Points
- Uniform interface: You can read from files, strings, network sockets, etc. using the same method.
- Error handling: nil means success, io.EOF means end of stream, other errors mean failure.
- Chunked reading: You decide how many bytes to read at a time by the size of p.

*/
// package main

// import (
// 	"fmt"
// 	"io"
// 	"strings"
// )

// func main(){
//     r:= strings.NewReader("Hello, Reader")
//     b:= make([]byte,8)

//     for{
//         n,err:= r.Read(b)
//         fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
//         		fmt.Printf("b[:n] = %q\n", b[:n])
// 		if err == io.EOF {
// 			break
// 		}
//     }
// }

/////////////////////////////////////////////////////////

package main 