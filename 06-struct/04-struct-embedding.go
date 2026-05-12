// Embed an Address struct inside Employee. Access the city field directly from the Employee instance,

package main

type Address struct {
	City    string
	Country string
}

type Employee struct {
	Name    string
	Address // This is struct embedding. Employee now has the fields of Address.
}

func main() {
	emp := Employee{
		Name: "Ayush S",
		Address: Address{
			City:    "Bangalore",
			Country: "India",
		},
	}

	// We can access the fields of the embedded struct directly from the Employee struct.
	println("Name:", emp.Name)
	println("City:", emp.City)       // Accessing City directly from Employee
	println("Country:", emp.Country) // Accessing Country directly from Employee
}
