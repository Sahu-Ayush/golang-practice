// Declare variables of different types

/*
Multiple variables of different data types can be declared under a single var keyword.
In the example shown below, since string b is not assigned any value,
a default empty string is assigned to it.
*/

package main

import "fmt"

func main() {
	var (
		a int = 10
		b string
		f float32 = -3.14
	)

	fmt.Println("a: ", a)
	fmt.Println("string b is not assigned any value: ", b)
	fmt.Println("f: ", f)
}
