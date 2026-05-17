package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func (u User) greet() {
	fmt.Println("Hello,", u.Name)
}

func main() {
	user := User{
		Name: "Ayush",
		Age:  25,
	}

	fmt.Println(user)
	fmt.Printf("%+v\n", user)
	fmt.Printf("%#v\n", user)
	fmt.Printf("%T\n", user)
	fmt.Printf("%p\n", &user)
	fmt.Println(user.Name)
	fmt.Println(user.Age)
	user.greet()
}