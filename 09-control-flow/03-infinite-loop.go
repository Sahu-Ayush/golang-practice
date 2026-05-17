package main

import "fmt"

func main() {
	i := 0
	for {
		// runs forever
    // use break to exit the loop
		fmt.Println("This will print indefinitely until we use break to exit the loop.")
		if i == 5 {
			break
		}
		i++
	}
}