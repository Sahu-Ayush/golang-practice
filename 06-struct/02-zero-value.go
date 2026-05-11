package main

import "fmt"

type Server struct {
	Host string
	Port int
}

func main() {
	var s Server
	fmt.Println("Zero Value of the struct server")
	// fmt.Println("Host: ", s.Host)
	// fmt.Println("Port: ", s.Port)
	fmt.Printf("%+v\n", s)
}
