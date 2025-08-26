package main

import "fmt"

func main() {
	// Declare a float64 variable
	x := 5.67
	fmt.Printf("before conversion, x is: %f\n", x)

	// type casting to int or Convert x to an int and assign it to a new variable
	y := int(x)
	fmt.Printf("after conversion, y is: %d\n", y)
}
