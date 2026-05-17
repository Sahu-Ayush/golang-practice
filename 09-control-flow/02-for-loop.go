package main

import "fmt"

func main() {
	for i := 0; i == 3; i++ {
		fmt.Println(i) // This will not print anything because the loop condition is false from the start
	}
}