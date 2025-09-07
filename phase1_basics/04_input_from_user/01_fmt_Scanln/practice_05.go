// Q5. Extra tokens left in buffer
// var x int
// var y int
// fmt.Scanln(&x)
// fmt.Scanln(&y)
// Input: 10 20 (in one line)
// What values go into x and y? Why?

package main

import "fmt"

func main() {
	var x int
	var y int

	fmt.Print("Enter two numbers: ")
	fmt.Scan(&x)
	fmt.Scan(&y)

	fmt.Printf("Scan one: %d\n", x)
	fmt.Printf("Scan two: %d\n", y)
}
