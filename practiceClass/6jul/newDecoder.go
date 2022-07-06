package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func main() {
	type x struct {
		A string
		B int
		C bool
	}

	jsonData := `{
		"A": "Hola Mundo",
		"B": 5,
		"C": true}`

	z := x{}

	buff := bytes.NewBuffer([]byte(jsonData))
	decoder := json.NewDecoder(buff)
	decoder.Decode(&z)

	fmt.Println(z)
}
