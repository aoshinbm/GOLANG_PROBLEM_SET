package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonInput := ` {
    "Thinly Sliced Peeled Apples": "6 cups",
    "sugar": "3/4 cup",
    "flour": "2 teaspoons",
    "cinnamon": "1/4 teaspoons",
    "nutmeg": "1/8 teaspoons"
    "lemon juice": "1 teaspoons"
    }`
	ingredients := make(map[string]string)

	err := json.Unmarshal([]byte(jsonInput), &ingredients)
	if err != nil {
		fmt.Println("JSON decode error!")
		return
	}
	fmt.Println(ingredients)
}
