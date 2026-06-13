package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
)

// Cursor is used for pagination.
type Cursor struct {
	Timestamp int64  `json:"ts"`
	ID        string `json:"id"`
}

// EncodeCursor encodes a cursor to a string.
func EncodeCursor(c Cursor) string {
	b, _ := json.Marshal(c)
	return base64.URLEncoding.EncodeToString(b)
}

// DecodeCursor decodes a cursor from a string.
// BUG: uses map[string]interface{} which causes JSON numbers to be parsed as float64,
// losing precision for large int64 values.
func DecodeCursor(s string) (Cursor, error) {
	b, err := base64.URLEncoding.DecodeString(s)
	if err != nil {
		return Cursor{}, err
	}
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return Cursor{}, err
	}
	ts, ok := raw["ts"].(float64)
	if !ok {
		return Cursor{}, fmt.Errorf("invalid timestamp type")
	}
	id, _ := raw["id"].(string)
	return Cursor{Timestamp: int64(ts), ID: id}, nil
}

func main() {
	c := Cursor{Timestamp: 1<<62, ID: "abc"}
	enc := EncodeCursor(c)
	dec, err := DecodeCursor(enc)
	if err != nil {
		panic(err)
	}
	fmt.Printf("original: %d\n", c.Timestamp)
	fmt.Printf("decoded:  %d\n", dec.Timestamp)
	if c.Timestamp != dec.Timestamp {
		fmt.Println("BUG: timestamp lost precision!")
	}
}
