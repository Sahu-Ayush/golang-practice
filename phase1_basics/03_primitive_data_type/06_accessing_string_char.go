package main

import "fmt"

func main() {

	myString := "Hello, Go!"
	fmt.Println(myString[7])        // prints: 71
	fmt.Printf("%c\n", myString[7]) // prints: G

}

/*
The character at index 7 is 'G'.

'G' in ASCII is 71, so myString[7] gives 71.

%c is used to print it as a character.
*/
