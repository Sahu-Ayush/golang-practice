package main

import "fmt"

func sum() (returnValue int) {
	defer func() {
		returnValue += 10
	}()
	return 5
}

func main() {
	result := sum()
	fmt.Println(result) // Output: 15
}

/*
In Go, when a function uses a named return value (like returnValue int in your sum() function), the deferred function can modify that named variable before the function actually returns to its caller.

Here's what happens step-by-step in your code:

sum() declares returnValue as a named return value, initialized to its zero value (0 for int).
The defer statement schedules an anonymous function to run after the main function body completes but before control returns to the caller.
The return 5 statement sets returnValue to 5.
The deferred function executes: returnValue += 10, changing it from 5 to 15.
The function returns 15, which is printed as the result.
This behavior is specific to named returns—deferred functions can't modify unnamed return values in the same way. If you changed the function to func sum() int (unnamed return), the defer wouldn't affect the output, and it would print 5 instead.
*/
