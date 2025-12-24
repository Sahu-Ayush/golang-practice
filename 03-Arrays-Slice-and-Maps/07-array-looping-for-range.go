// Array looping using for loop with range keyword
package main

import (
	"fmt"
)

func main() {
	var grades [5]int = [5]int{10, 20, 30, 40, 50}

	for index, element := range grades {
		fmt.Println("grades[", index, "] value is", element)
	}
}

/* Output:
grades[ 0 ] value is 10
grades[ 1 ] value is 20
grades[ 2 ] value is 30
grades[ 3 ] value is 40
grades[ 4 ] value is 50
*/
