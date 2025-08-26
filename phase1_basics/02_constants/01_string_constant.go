// String constants

/*
They have their values enclosed within double-quotes. Typed string constants have the data type string included within their syntax.
In contrast, untyped constants have their data types identified by Type inference.
The example below declares an untyped string constant whose value is then assigned to a variable. The word variable acquires its data type from the string constant.
Constant s is a typed string constant.
*/

package main

import "fmt"

func main() {
	const n = "Hello"
	var word = n

	fmt.Printf("Type %T value %v\n", word, word)

	const s string = "World"
	fmt.Printf("Type %T value %v\n", s, s)

}
