// Data Types in Go (Primitive Types)

package main

import (
	"fmt"
	"reflect"
)

func main() {
	var age int = 28
	var height float64 = 5.9
	var name string = "Ayush"
	var isCoder bool = true

	fmt.Println("Age:", age, "| Type:", reflect.TypeOf(age))
	fmt.Println("Height:", height, "| Type:", reflect.TypeOf(height))
	fmt.Println("Name:", name, "| Type:", reflect.TypeOf(name))
	fmt.Println("Is Coder:", isCoder, "| Type:", reflect.TypeOf(isCoder))
}
