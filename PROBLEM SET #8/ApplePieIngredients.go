package main

import (
	"encoding/json"
	"fmt"
	"reflect"
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
	// Declared an empty map
	ingredients := make(map[string]string)

	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal([]byte(jsonInput), &ingredients)
	//if err != nil {
	//	fmt.Println("JSON decode error!")
	//	return
	//}
	// Print the data type
	fmt.Println(reflect.TypeOf(ingredients))

}
