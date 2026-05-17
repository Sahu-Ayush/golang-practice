package main

import "fmt"

type Celsius float64

func (c Celsius) toFahrenheit() float64 {
	return float64(c)*9/5 + 32
}

func main() {
	temp := Celsius(25)

	fmt.Println(temp.toFahrenheit())
}