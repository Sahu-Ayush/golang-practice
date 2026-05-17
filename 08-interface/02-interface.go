package main

import "fmt"

// Define an interface — a set of method signatures
type Shape interface {
    Area() float64
    Perimeter() float64
}

// Concrete type: Circle
type Circle struct {
    Radius float64
}

func (c Circle) Area() float64      { return 3.14 * c.Radius * c.Radius }
func (c Circle) Perimeter() float64 { return 2 * 3.14 * c.Radius }

// Concrete type: Rectangle
type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64      { return r.Width * r.Height }
func (r Rectangle) Perimeter() float64 { return 2 * (r.Width + r.Height) }

// Function that works with ANY Shape
func PrintInfo(s Shape) {
    fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

func main() {
    // PrintInfo(Circle{Radius: 5})
    // PrintInfo(Rectangle{Width: 4, Height: 6})

		c := Circle{Radius: 5}
		r := Rectangle{Width: 4, Height: 6}

		fmt.Println("Circle:")
		PrintInfo(c)

		fmt.Println("Rectangle:")
		PrintInfo(r)		
}