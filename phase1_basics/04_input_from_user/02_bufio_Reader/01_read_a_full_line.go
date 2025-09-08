// Reading a full line using bufio.Reader

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n') // Reads until Enter
	name = strings.TrimSpace(name)

	fmt.Print("Enter your city: ")
	city, _ := reader.ReadString('\n')
	city = strings.TrimSpace(city) // remove space and newline

	fmt.Printf("Hello, %s from %s!\n", name, city)

}
