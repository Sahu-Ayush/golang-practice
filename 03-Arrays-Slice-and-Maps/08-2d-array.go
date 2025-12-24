// Multidimensional (2D) arrays and accessing specific elements in Go
package main

import "fmt"

func main() {
	// Declare and initialize a 2D array
	var arr [3][2]int = [3][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}

	// Access and print specific elements
	fmt.Println("Element at row 0, column 1:", arr[0][1])
	fmt.Println("Element at row 1, column 0:", arr[1][0])
	fmt.Println("Element at row 2, column 1:", arr[2][1])

}

/* Output:
Element at row 0, column 1: 2
Element at row 1, column 0: 3
Element at row 2, column 1: 6
*/
