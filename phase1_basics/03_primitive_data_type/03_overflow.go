package main

import "fmt"

func main() {

	// initializing with the maximum allowed negative decimal value
	var num float64 = -1.7e+308

	// printing the value and data type
	fmt.Printf("Value is: %f and type is: %T\n", num, num)

	// making the value out of range by multiply by +/-10
	var num_n float64 = num * 10
	var num_p float64 = num * -10

	// printing out new value and type
	fmt.Printf("Value is: %f and type is: %T\n", num_n, num_n)
	fmt.Printf("Value is: %f and type is: %T\n", num_p, num_p)

}
