// Update an element in an array by its index.
// Access an element of an array by its index
package main

import (
	"fmt"
)

func main() {
	var grades [5]int = [5]int{10, 20, 30, 40, 50}
	fmt.Println("grades: ", grades)

	grades[2] = 100
	fmt.Println("Updated grades: ", grades)
}

/* Output:
grades:  [10 20 30 40 50]
Updated grades:  [10 20 100 40 50]
*/
