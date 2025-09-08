// Reading multiple lines using bufio.Scanner

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter 3 lines of text: ")
	for i := 1; i <= 3; i++ {
		scanner.Scan() // waits for input
		text := scanner.Text()
		fmt.Printf("Line: %d: %s\n", i, text)
	}
}
