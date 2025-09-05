// Multiple variables in one line

// Write a program that reads two strings (first name, last name) using one Scanln.

// Input: Ayush Sahu
// Output: Hello Ayush Sahu

package main

import "fmt"

func main() {
	var firstName, lastName string

	fmt.Print("Enter your full name: ")
	n, err := fmt.Scanln(&firstName, &lastName)

	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Printf("Hello %s %s", firstName, lastName)
		fmt.Printf("Total token %d\n", n)
	}
}
