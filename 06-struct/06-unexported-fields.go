package main

import "fmt"

type Counter struct {
	count int // unexported field
}

// Increment increases the counter by one.
// This method can access the unexported field `count` because it's defined within the same package.
// However, code outside this package cannot access `count` directly, ensuring encapsulation.
// This allows us to control how the counter is modified and prevents external code from manipulating it in unintended ways.
// By providing methods like `Increment`, we can safely modify the internal state of the `Counter` without exposing its implementation details.
// This design promotes better encapsulation and allows us to maintain invariants about the `Counter`'s state.
// The `Increment` method is a public method that can be called from outside the package, while the `count` field remains private and inaccessible to external code.
// This is a common pattern in Go to provide controlled access to internal state while keeping the implementation details hidden.
// In this example, the `Counter` struct has an unexported field `count`, which can only be accessed and modified through the `Increment` method. This allows us to maintain control over how the counter is updated while keeping the internal state hidden from external code.

// HOW c *Counter works: When we call `Increment` on a `Counter` instance, we pass a pointer to that instance (`*Counter`). This allows the method to modify the original `Counter` struct's state, specifically the `count` field. If we were to use a value receiver (i.e., `func (c Counter) Increment()`), it would operate on a copy of the `Counter`, and any modifications to `count` would not affect the original instance. By using a pointer receiver, we ensure that changes made within the `Increment` method are reflected in the original `Counter` instance.
func (c *Counter) Increment() {
	c.count++
}

func (c *Counter) Value() int {
	return c.count
}

func main() {
	counter := &Counter{} // Create a new Counter instance
	counter.Increment()   // Increment the counter
	counter.Increment()   // Increment the counter again

	val := counter.Value() // Get the current value of the counter
	fmt.Println(val)       // Output: 2
}
