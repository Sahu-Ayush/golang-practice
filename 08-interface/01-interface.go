package main

import (
	"fmt"
	"reflect"
)

type Speaker interface {
	speak()
}

type Dog struct {
	Name string
}

func (d Dog) speak() {
	fmt.Println("[Dog.speak()] Method execution started")

	fmt.Println("[Dog.speak()] Receiver value:", d)

	fmt.Println("[Dog.speak()] Dog says: Woof")

	fmt.Println("[Dog.speak()] Method execution completed")
}

func makeSpeak(s Speaker) {
	fmt.Println("--------------------------------------------------")
	fmt.Println("[makeSpeak()] Function started")

	// Interface actual type
	fmt.Println("[makeSpeak()] Interface contains concrete type:")
	fmt.Println("   ", reflect.TypeOf(s))

	// Interface actual value
	fmt.Println("[makeSpeak()] Interface contains concrete value:")
	fmt.Println("   ", s)

	fmt.Println("[makeSpeak()] Calling s.speak()")

	// Dynamic dispatch happens here
	s.speak()

	fmt.Println("[makeSpeak()] Returned from s.speak()")

	fmt.Println("[makeSpeak()] Function completed")
	fmt.Println("--------------------------------------------------")
}

func main() {
	fmt.Println("[main()] Program started")

	fmt.Println()

	fmt.Println("[main()] Creating Dog object")

	d := Dog{
		Name: "Rocky",
	}

	fmt.Println("[main()] Dog object created:")
	fmt.Println("   ", d)

	fmt.Println()

	fmt.Println("[main()] Passing Dog into makeSpeak(d)")
	fmt.Println("[main()] Compiler checks:")
	fmt.Println("   Does Dog implement Speaker?")
	fmt.Println("   Dog has speak() method")
	fmt.Println("   YES -> Dog satisfies Speaker")

	fmt.Println()

	// Interface conversion happens here
	makeSpeak(d)

	fmt.Println()

	fmt.Println("[main()] Program completed")
}