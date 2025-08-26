package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Simple input using fmt.Scanln
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	fmt.Println("Hello,", name)
	// fmt.Println()

	// Buffered input using bufio
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your age: ")
	ageStr, _ := reader.ReadString('\n') // includes newline
	ageStr = strings.TrimSpace(ageStr)   // remove newline and other whitespace

	// Error handling for strconv.Atoi
	age, err := strconv.Atoi(ageStr) // convert to int
	if err != nil {
		fmt.Println("\nInvalid age entered. Please enter a number.")
		// Optionally, you can exit or ask for input again
		return
	}

	fmt.Println("You are", age, "years old.")
}
