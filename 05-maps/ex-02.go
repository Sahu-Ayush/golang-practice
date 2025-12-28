package main

import "fmt"

func main() {
	var ascii_codes map[string]int
	ascii_codes = map[string]int{} // initialize the map
	// or use: ascii_codes := make(map[string]int)
	// or: ascii_codes := map[string]int{} to both declare and initialize
	ascii_codes["A"] = 65
	fmt.Println(ascii_codes)
}
