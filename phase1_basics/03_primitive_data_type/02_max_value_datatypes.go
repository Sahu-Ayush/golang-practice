// To find the maximum value of a specific data type

/*
The math package in Go provides constants representing the maximum (and minimum for signed types) values for various integer sizes.
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	// For Integer Types (signed and unsigned)

	// Signed integers
	fmt.Printf("Max int8: %d, Min int8: %d\n", math.MaxInt8, math.MinInt8)
	fmt.Printf("Max int16: %d, Min int16: %d\n", math.MaxInt16, math.MinInt16)
	fmt.Printf("Max int32: %d, Min int32: %d\n", math.MaxInt32, math.MinInt32)
	fmt.Printf("Max int64: %d, Min int64: %d\n", math.MaxInt64, math.MinInt64)

	// Unsigned integers
	fmt.Printf("Max uint8: %d, Min uint8: %d\n", math.MaxUint8, 0)

	// Platform-dependent int and unint
	fmt.Printf("Max int: %d, Min int: %d\n", math.MaxInt, math.MinInt)
	fmt.Printf("Max uint: %d, Min uint: %d\n", uint(math.MaxUint), 0)

	// For Floating-Point Types:
	fmt.Printf("Max float32: %e, Min float32: %e\n", math.MaxFloat32, math.SmallestNonzeroFloat32)
	fmt.Printf("Max float64: %e, Min float64: %e\n", math.MaxFloat64, math.SmallestNonzeroFloat64)

}
