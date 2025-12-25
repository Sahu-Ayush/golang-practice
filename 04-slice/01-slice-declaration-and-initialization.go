// Slice Declaration and Initialization in Go
package main

import (
	"fmt"
)

func main() {
	// Declaring and initializing a slice
	grades := []int{10, 20, 30, 40, 50}
	fmt.Println(grades)

	// declaring a slice without initialization
	var numbers []int
	fmt.Println(numbers)

	// Initializing a slice
	numbers = []int{1, 2, 3, 4, 5}
	fmt.Println(numbers)
}
