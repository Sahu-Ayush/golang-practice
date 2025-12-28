// map without initializer and len function

package main

import "fmt"

func main() {
	var m map[string]int
	fmt.Println("Length of map m:", len(m)) // Output: Length of map m: 0
}
