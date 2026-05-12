//Add JSON tags to a Config struct so it marshals to snake_case keys. Marshal it and print the JSON.

package main

import (
	"encoding/json"
	"fmt"
)

type Config struct {
	MaxRetries int    `json:"max_retries"`
	TimeoutSec int    `json:"timeout_sec"`
	BaseURL    string `json:"base_url"`
}

func main() {
	data := Config{
		MaxRetries: 5,
		TimeoutSec: 30,
		BaseURL:    "https://api.example.com",
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// WHY string(jsonData) instead of just jsonData?
	// jsonData is a byte slice ([]byte), and we need to convert it to a string to print it in a human-readable format.
	// If we print jsonData directly, it will show the byte slice representation, which is not what we want.
	// By using string(jsonData), we convert the byte slice to a string, allowing us to see the actual JSON content.
	fmt.Println(string(jsonData))
}
