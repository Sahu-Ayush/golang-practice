// Complex Constants

/*
Complex constants are represented as an ordered pair of an integer constant and an imaginary constant.
The example shows number 5  as a complex number because of the data type assigned to it.
*/

package main

import "fmt"

func main() {
	const ch complex64 = 5

	fmt.Printf("Type %T value %v\n", ch, ch)
}
