package main

import "fmt"

const pi = 3.14159           // Untyped constant
const language string = "Go" // Typed constant

func main() {
	const daysInWeek = 7 // Constant integer
	const isGoFun = true // Constant boolean

	fmt.Println("Pi:", pi)
	fmt.Println("Language:", language)
	fmt.Println("Days in a week:", daysInWeek)
	fmt.Println("Is Go fun?", isGoFun)
}
