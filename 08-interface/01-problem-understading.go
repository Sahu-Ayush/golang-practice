// Animal Sounds

// WITHOUT Interface

package main

import "fmt"

// type Dog struct {
// }

// func (d Dog) speak() {
// 	fmt.Println("Woof!")
// }

// type Cat struct {
// }

// func (c Cat) speak() {
// 	fmt.Println("Meow!")
// }

// func main() {
// 	dog := Dog{}
// 	cat := Cat{}

// 	dog.speak()
// 	cat.speak()
// }


// ---

// type Dog struct {
// }

// func (d Dog) speak() {
// 	fmt.Println("Woof!")
// }

// type Cat struct {
// }

// func (c Cat) speak() {
// 	fmt.Println("Meow!")
// }

// func MakeDogSpeak(d Dog){
// 	fmt.Println("The dog says: ")
// 	d.speak()
// }

// func MakeCatSpeak(c Cat){
// 	fmt.Println("The cat says: ")
// 	c.speak()
// }

// func main() {
// 	dog := Dog{}
// 	cat := Cat{}

// 	MakeDogSpeak(dog)
// 	MakeCatSpeak(cat)
// }

// type Animal interface {
//     Dog | Cat
//     speak()
// }

type Dog struct {
	Name string
}

func (d Dog) speak() {
	fmt.Println("Woof!")
}

type Cat struct {
	Name string
}

func (c Cat) speak() {
	fmt.Println("Meow!")
}

// func MakeAnimalSpeak[T Animal](s T) {
//     fmt.Println("The animal says:")
//     s.speak()
// }
func MakeAnimalSpeak[T interface{ Dog | Cat; speak() }](s T) {
	fmt.Println("The animal says: ")
	s.speak()
}


func main() {
	dog := Dog{Name: "Buddy"}
	cat := Cat{Name: "Whiskers"}

	MakeAnimalSpeak(dog)
	MakeAnimalSpeak(cat)
}