// Declare and Initialize a Slice Using Another Slice
package main

import (
	"fmt"
)

func main() {
	// Declaring and initializing a slice
	originalSlice := []string{"Go", "is", "an", "open-source", "programming", "language"}
	fmt.Println("Original Slice:", originalSlice)

	// Declaring and initializing a new slice using the original slice
	newSlice := originalSlice[1:4] // Slicing from index 1 to 4 (4 is exclusive)
	fmt.Println("New Slice from Original Slice (index 1 to 4):", newSlice)

	// Another example of slicing
	anotherSlice := originalSlice[:3] // Slicing from start to index 3 (3 is exclusive)
	fmt.Println("Another Slice from Original Slice (start to index 3):", anotherSlice)
}

/* Output:
Original Slice: [Go is an open-source programming language]
New Slice from Original Slice (index 1 to 4): [is an open-source]
Another Slice from Original Slice (start to index 3): [Go is an]
*/
