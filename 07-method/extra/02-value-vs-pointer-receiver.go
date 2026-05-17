package main

import (
    "fmt"
)

type Counter struct {
    value int
}

func (c Counter) IncrementByOne() int {
    c.value++
    fmt.Println("[Increment]: Counter value inside method:", c.value, "after IncrementByOne using value receiver")
    return c.value
}

func (c *Counter) IncrementByTwo() int {
    c.value += 2
    fmt.Println("[Increment]: Counter value inside method:", c.value, "after IncrementByTwo using pointer receiver")
    return c.value
}

func main() {
    // Question: does c1 or c2 store an address?
    // Answer:
    //   - c1 is a value variable of type Counter.
    //   - c2 is a pointer variable of type *Counter.
    // The proof uses fmt.Printf to show the actual types and addresses.

    c1 := Counter{value: 0}
    // c1 is a Counter value. Printing its type shows "main.Counter".
    // &c1 is the address where c1 is stored in memory.
    fmt.Printf("c1 type=%T, c1=%v, &c1=%p\n", c1, c1, &c1)
    fmt.Println("[main]: Counter value before calling IncrementByOne():", c1.value)
    c1.IncrementByOne()
    fmt.Println("[main]: Counter value after calling IncrementByOne():", c1.value)
    c1.IncrementByTwo()
    fmt.Println("[main]: Counter value after calling IncrementByTwo():", c1.value)

    fmt.Println("--------------------------------------------------")

    c2 := &Counter{value: 10}
    // c2 is a pointer to a Counter. Printing its type shows "*main.Counter".
    // c2 itself is an address (pointer) to the Counter value.
    // &c2 is the address of the pointer variable c2.
    fmt.Printf("c2 type=%T, c2=%v, c2=%p, &c2=%p\n", c2, c2, c2, &c2)
    fmt.Println("[main]: Counter value before calling IncrementByOne():", c2.value)
    c2.IncrementByOne()
    fmt.Println("[main]: Counter value after calling IncrementByOne():", c2.value)
    c2.IncrementByTwo()
    fmt.Println("[main]: Counter value after calling IncrementByTwo():", c2.value)
}
