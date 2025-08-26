// Complex Constants

/*
They are very much similar to string constants but can define only true or false.
*/

package main

import "fmt"

func main() {
	const a bool = true
	fmt.Printf("Type %T value %v\n", a, a)
}
