package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		i := i // Capture the current value of i
		defer fmt.Println(i)
	}
}
