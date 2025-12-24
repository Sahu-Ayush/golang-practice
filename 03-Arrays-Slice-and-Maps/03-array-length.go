// To Find length of an Array in Go using len() function
package main

import (
	"fmt"
)

func main() {
	// initialize array
	var arr = [5]int{10, 20, 30, 40, 50}

	// find length of array
	var length int
	length = len(arr)

	// print length of array
	fmt.Println("Length of array is:", length)
}
