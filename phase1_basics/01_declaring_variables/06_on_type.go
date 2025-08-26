// Variable Declaration with no type or Type Inference

/*
var <varibale_name> = <value>
*/

package main

import "fmt"

func main() {
	var a = 10
	var f = -3.14

	fmt.Printf("%v is of %T data type\n", a, a)
	fmt.Printf("%v is of %T data type", f, f)
}
