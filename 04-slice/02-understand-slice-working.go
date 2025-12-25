// Understanding how Slice Work in Go
package main

import (
	"fmt"
)

func main() {
	// Declaring and initializing a slice mentioning "lower and upper" bounds or "start and end" indexs
	arr := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Println("Original Array:", arr)

	// Slicing the array from index 0 to 3 (3 is exclusive)
	slice1 := arr[0:3]
	fmt.Println("Slice from index 0 to 3:", slice1)

	// Slicing the array from index 1 to 10 (10 is exclusive) (till last element)
	slice2 := arr[1:10]
	fmt.Println("Slice from index 1 to 10:", slice2)

	// Do not mension lower bound (start index), it will consider 0 as default
	slice3 := arr[:4]
	fmt.Println("Slice from start to index 4:", slice3)

	// Do not mension upper bound (end index), it will consider len(arr) as default
	slice4 := arr[5:]
	fmt.Println("Slice from index 5 to end:", slice4)

	// Do not mension both lower and upper bounds, it will consider the whole array
	slice5 := arr[:]
	fmt.Println("Slice of the whole array:", slice5)
}

/* Output:
Original Array: [10 20 30 40 50 60 70 80 90 100]
Slice from index 0 to 3: [10 20 30]
Slice from index 1 to 11: [20 30 40 50 60 70 80 90 100]
Slice from start to index 4: [10 20 30 40]
Slice from index 5 to end: [60 70 80 90 100]
Slice of the whole array: [10 20 30 40 50 60 70 80 90 100]
*/
