package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	var p Person
	p = Person{
		Name: "Ayush",
		Age:  22,
	}
	fmt.Println(p)
}

// Define a struct called "Person" with Name (string) and Age (int). Create one instance and print both fields.
//  Output:{Ayush 22}
