package main

import "fmt"

// interface: Generic stack interface
type Stack[T any] interface {
	Push(n T)
	Pop() T
}

// specific implementation
type intStack struct {
	items []int
}

func (i *intStack) Pop() int {
	if len(i.items) == 0 {
		return 0
	}
	item := i.items[len(i.items)-1]
	i.items = i.items[:len(i.items)-1]
	return item
}

func (i *intStack) Push(n int) {
	i.items = append(i.items, n)
}

// generic implementation
type tStack[T any] struct {
	items []T
}

func (t *tStack[T]) Push(n T) {
	t.items = append(t.items, n)
}

func (t *tStack[T]) Pop() T {
	if len(t.items) == 0 {
		var zero T
		return zero
	}
	item := t.items[len(t.items)-1]
	t.items = t.items[:len(t.items)-1]
	return item
}

// Polymorphic function that accepts ANY integer stack
func UseIntStack(s Stack[int]) {
	s.Push(100)
	s.Push(200)
	fmt.Println(s)
	fmt.Println("popped from stack: ", s.Pop())
}

// for by passing the default formatting:
func (i *intStack) String() string {
	return fmt.Sprintf("%v", i.items)
}

func (t *tStack[T]) String() string {
	return fmt.Sprintf("%v", t.items)
}

/*
How fmt.Println Uses String() string Behind the Scenes
In the standard library fmt package, Go defines the Stringer interface:
type Stringer interface {
    String() string
}
Whenever you pass any variable to fmt.Println, fmt.Printf("%v"), or fmt.Sprint, the fmt package checks if that variable implements fmt.Stringer using a runtime type assertion equivalent to this:
if stringer, ok := value.(fmt.Stringer); ok {
    return stringer.String()
}
If your struct (or pointer receiver) implements String() string:

1. Go detects that it satisfies fmt.Stringer.

2. It bypasses the default struct formatting (&{...}).

3. It calls your String() method directly and prints whatever string you return.
*/
func main() {

	inSt := &intStack{}
	tSt := &tStack[int]{}

	fmt.Println("Using intStack: ")
	UseIntStack(inSt)
	fmt.Println("Using tStack[int]: ")
	UseIntStack(tSt)

}
