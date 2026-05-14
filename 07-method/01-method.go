package main

import "fmt"

// Methods are functions with a receiver argument, allowing us to associate behavior with structs.
type Counter struct {
    Value int
    Name  string
}

// Value receiver — works on a COPY, doesn't mutate
func (c Counter) Display() string {
    return fmt.Sprintf("%s: %d", c.Name, c.Value)
}

// Pointer receiver — mutates the original struct
func (c *Counter) Increment() {
    c.Value++   // modifies the actual struct
}

func (c *Counter) Reset() {
    c.Value = 0
}

func main() {
    c := &Counter{Name: "requests"}
    c.Increment()
    c.Increment()
    fmt.Println(c.Display()) // requests: 2
    c.Reset()
    fmt.Println(c.Display()) // requests: 0
}