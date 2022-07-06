package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	type product struct {
		Name      string
		Price     int
		Published bool
	}

	jsonData := `{"Name": "MacBook Air", "Price": 900, "Published": true}`

	var p product

	if err := json.Unmarshal([]byte(jsonData), &p); err != nil {
		log.Fatal(err)
	}

	fmt.Println(p)
}
