package main

import "fmt"

func main() {
	var smallNum int8 = 2
	fmt.Println("Small Number:", smallNum)
	// var overflowNum int8 = 128
	// fmt.Println("Small Number:", overflowNum)
	/*
		Note:
		error: cannot use 128 (untyped int constant) as int8 value in variable declaration (overflows)compilerNumericOverflow
		doc: https://pkg.go.dev/golang.org/x/tools/internal/typesinternal#NumericOverflow
	*/
}
