package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		if i == 3 {
			continue // Skip the rest of the loop body when i is 3
		}
		if i == 7 {
			break // Exit the loop when i is 7
		}
		fmt.Println(i)
	}
}