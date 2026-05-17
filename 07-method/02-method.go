// This example demonstrates the difference between value and pointer receivers in Go methods, and how they interact with addressability of values.
// In Go, methods can have either value receivers (e.g., `func (c Counter)`) or pointer receivers (e.g., `func (c *Counter)`). The choice between them affects how the method can modify the struct and whether it can be called on non-addressable values (like temporary structs returned from functions).
// In this code, we have a `Counter` struct with a `Value` field. We define two methods: `Increment` with a pointer receiver that modifies the original struct, and `Display` with a value receiver that returns a string representation of the counter's value.
// The `NewCounter` function is a constructor that returns a new `Counter` instance with an initial value of 0.
// In the `main` function, we demonstrate that calling `Display` on a non-addressable value (the result of `NewCounter()`) works fine because it uses a value receiver, which operates on a copy of the struct. However, calling `Increment` on the same non-addressable value results in a compile-time error because it requires a pointer receiver, which cannot be called on a temporary value.
// To successfully call `Increment`, we need to store the result of `NewCounter()` in a variable (making it addressable) and then call `Increment` on that variable. Go will automatically take the address of the variable when calling the pointer method, allowing us to modify the original struct.	
package main

import "fmt"

type Counter struct {
    Value int
}

func (c *Counter) Increment() {
    c.Value++
}

func (c Counter) Display() string {
    return fmt.Sprintf("Value: %d", c.Value)
}

func NewCounter() Counter {
    return Counter{Value: 2}
}

func main() {
    // ✅ Value receiver works on non-addressable values
    // NewCounter().Display()  // OK - works fine, value is copied
		// fmt.Println(NewCounter().Display())

		n := NewCounter()
		fmt.Println(n.Display())
    // // ❌ Pointer receiver fails on non-addressable values
    // NewCounter().Increment()  // ERROR: cannot call pointer method on non-addressable Counter
    
    // ✅ But pointer receiver works if value is addressable (stored in variable)
    c := NewCounter()
    c.Increment()  // OK - c is addressable, Go auto-addresses it to &c

		fmt.Println("After Increment:", c.Display())
}