// Typed and Untyped Numeric constants

/*
Typed Numeric constants are declared with their data type. Once declared they cannot hold values of any other data type.
Untyped Numeric constants on the other hand are declared without mentioning their data type. The data type is identified by Type inference.
Constant num1 is untyped while num2 is typed.
*/

/*
Numeric constants can be assigned integers, floating-point numbers and complex numbers.
An Integer constant is assigned integer values including decimal, octal and hexadecimal.
*/

/*

 */

package main

import "fmt"

func main() {
	const num1 = 2.3

	fmt.Printf("Type %T value %v\n", num1, num1)

	const num2 int = 4

	fmt.Printf("Type %T value %v\n", num2, num2)

	const num = 5
	var a = num
	fmt.Printf("Type %T value %v\n", a, a)

}
