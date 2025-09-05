// Single word input
// Write a program that asks for your favorite color and prints it back.
// Input: Blue
// Output: Your favorite color is Blue.

package main

import "fmt"

func main() {
	var color string

	fmt.Print("Enter your favorite color: ")
	fmt.Scanln(&color)

	fmt.Printf("Your favorite color is %s\n", color)
}
