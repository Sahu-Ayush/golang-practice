// Short variable declaration

/*
The := operator in Go is used to declare variables without mentioning the var keyword and data type.
Like the previous case,
Type Inference is used to identify the data type of the value assigned to the variable.
*/

package main

import "fmt"

func main() {
	a := 10
	f := -3.14

	fmt.Printf("%v is of %T data type\n", a, a)
	fmt.Printf("%v is of %T data type", f, f)
}
