package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	// Named variable + struct literal initialization.
	// Use this when you need to reuse the same Person value multiple times.
	var p Person
	p = Person{
		Name: "Ayush",
		Age:  22,
	}
	// fmt.Println(p)
	fmt.Println(p.Name, p.Age)

	// Print the whole struct value with a formatting verb.
	fmt.Printf("%+v\n", p) // Output: {Name:Ayush Age:22}

	// Directly using a struct literal and accessing its field.
	// This is fine for a one-off value when you don't need to store it.
	fmt.Println((Person{
		Name: "Ayush",
		Age:  22,
	}).Name)

	// If you want to use fmt.Printf, pass a format string explicitly.
	fmt.Printf("%s\n", (Person{
		Name: "Ayush",
		Age:  22,
	}).Name)
}

// Define a struct called "Person" with Name (string) and Age (int). Create one instance and print both fields.
// Output:
// Ayush 22
// {Name:Ayush Age:22}
// Ayush
// Ayush