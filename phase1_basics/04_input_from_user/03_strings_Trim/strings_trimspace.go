package main

import (
	"fmt"
	"strings"
)

func main() {
	input1 := "   Ayush   "
	fmt.Printf("[%s]\n", strings.TrimSpace(input1))

	input2 := "Ayush\n"
	fmt.Printf("[%s]\n", strings.TrimSpace(input2))

	input3 := "Ayush   Sahu"
	fmt.Printf("[%s]\n", strings.TrimSpace(input3))

	input4 := "\t  Ayush Sahu \n"
	fmt.Printf("[%s]\n", strings.TrimSpace(input4))
}

// Case 1: spaces before and after
// [Ayush]

// Case 2: newline at the end
// [Ayush]

// Case 3: spaces inside (not removed!)
// [Ayush   Sahu]

// Case 4: tabs and newlines
// [Ayush Sahu]
