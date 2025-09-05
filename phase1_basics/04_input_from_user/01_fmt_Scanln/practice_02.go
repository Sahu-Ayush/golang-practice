// Multiple word issue

// Try this program:

package main

import "fmt"

func main() {

	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	fmt.Printf("Hello %s!\n", name)
}

// Input: Ayush Sahu

// What gets stored in name? Why?

// "name" variable store "Ayush" token as Scanln stop reading at first space or newline and "Sahu" second token will goes under stdin buffer
