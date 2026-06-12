package main

import (
	"encoding/json"
	"fmt"
)

// TODO: define a struct that can distinguish between missing, null, and present fields.

func main() {
	// Case 1: "age" is present and set to 30
	case1 := []byte(`{ "name": "Alice", "age": 30 }`)

	// Case 2: "age" is explicitly null
	case2 := []byte(`{ "name": "Bob", "age": null }`)

	// Case 3: "age" is missing entirely
	case3 := []byte(`{ "name": "Charlie" }`)

	for _, raw := range [][]byte{case1, case2, case3} {
		var person Person
		if err := json.Unmarshal(raw, &person); err != nil {
			fmt.Println("error:", err)
			continue
		}
		fmt.Printf("%+v\n", person)
	}
}
