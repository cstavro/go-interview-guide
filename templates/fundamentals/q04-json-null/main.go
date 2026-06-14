package main

import (
	"encoding/json"
	"fmt"
)

// Person is intentionally empty. The candidate must add fields
// to distinguish between missing, null, and present values.
type Person struct{}

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
