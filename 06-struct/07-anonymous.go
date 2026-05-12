// Use an anonymos struct to decode a one-off JSN response without defining a name type.

package main

import (
	"encoding/json"
	"fmt"
)

// func main() {
// 	raw := []bytrs('{"status"}: "Ok", "code": 200}')

// 	// Declare an anonymous struct type. The `var` declaration here
// 	// creates `response` with the struct's zero value automatically.
// 	// This is the same as using an explicit composite literal like:
// 	//   var response struct { Status string `json:"status"`; Code int `json:"code"` }{}
// 	var response struct {
// 		Status string `json:"status"`
// 		Code   int    `json:"code"`
// 	}

// Use an explicit named struct type to decode a JSON response.

type Response struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
}

func main() {
	raw := []byte(`{"status": "Ok", "code": 200}`)

	// Declare a named type and use it explicitly.
	// `var response Response` creates the zero value for the `Response` struct.
	var response Response
	err := json.Unmarshal(raw, &response)
	if err != nil {
		panic(err)
	}

	fmt.Println(response)

}
