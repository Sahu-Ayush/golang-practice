// Int to Float Program

package main

import "fmt"

func main() {
	// Declare an int variable
	x := 5
	fmt.Printf("before conversion, x is: %d\n", x)

	// type casting to flat64 or Convert x to a float64 and assign it to a new variable
	y := float64(x)
	fmt.Printf("after conversion, y is: %f\n", y)
}
