package main

import (
	"fmt"
)

func main() {
	// Array declaration
	var grades [5]int
	fmt.Println("grades:", grades)

	var fruits [3]string
	fmt.Println("fruits:", fruits)

	// Array initialization
	var scores = [4]int{90, 80, 85, 95}
	fmt.Println("scores:", scores)

	names := [3]string{"Alice", "Bob", "Charlie"}
	fmt.Println("names:", names)

	// Accessing and modifying array elements
	scores[2] = 88
	fmt.Println("Updated scores:", scores)

	// Array length
	fmt.Println("Length of names array:", len(names))
}

/* Expected Output:
grades: [0 0 0 0 0]
fruits: [  ]
scores: [90 80 85 95]
names: [Alice Bob Charlie]
Updated scores: [90 80 88 95]
Length of names array: 3
*/
