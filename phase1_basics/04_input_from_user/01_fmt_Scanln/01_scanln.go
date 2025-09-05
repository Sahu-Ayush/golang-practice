// Reading integers and strings

package main

import (
	"fmt"
)

func main() {
	var name string
	var age int

	fmt.Print("Enter your name: ")
	fmt.Scanln(&name) // reads until first space or newline

	fmt.Print("Enter your age: ")
	fmt.Scanln(&age) // reads integer

	fmt.Printf("Hello, %s! You are %d years old.", name, age)
}
