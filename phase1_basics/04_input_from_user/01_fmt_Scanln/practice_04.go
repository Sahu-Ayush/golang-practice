// Q4. Wrong type entered

// Ask the user for their age (int).

// Input: twenty

// What happens to the variable? What error is returned?

package main

import "fmt"

func main() {
	var age int

	fmt.Print("Enter you age: ")
	n, err := fmt.Scanln(&age)

	fmt.Printf("Scanned: %d\n", n)
	if err != nil {
		fmt.Print(err.Error(), "\n")
	} else {
		fmt.Printf("Your age is %d\n", age)
	}
}

// Ans: expected integer
