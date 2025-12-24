package main

import "fmt"

func main() {
	// Array initialization using curly braces
	var grades [3]int = [3]int{10, 20, 30}
	fmt.Println("grades: ", grades)

	// Array initialization using Shorthand declaration
	grades_2 := [4]int{10, 20, 30, 40}
	fmt.Println("grades_2: ", grades_2)

	// Array initialization with ellipsis
	// The size of the array is determined by the number of elements provided
	grades_3 := [...]int{10, 20, 30, 40, 50}
	fmt.Println("grades_3: ", grades_3)
}

/* Expected Output:
grades:  [10 20 30]
grades_2:  [10 20 30 40]
grades_3:  [10 20 30 40 50]
*/
