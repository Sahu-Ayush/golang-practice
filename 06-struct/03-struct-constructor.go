// Write a NewServer(host string, port int) constructor that returns a *Server with a default timeout of 30.

package main

import "fmt"

type Server struct {
	Host		string
	Port		int
	Timeout		int
}

func NewServer(host string, port int) *Server {
	return &Server {
		Host: host,
		Port: port,
		Timeout: 30,
	}
}

func main() {
	ser := NewServer("localhost", 8080)
	fmt.Printf("%+v\n", ser)
}

// Output:
// &{Host:localhost Port:8080 Timeout:30}