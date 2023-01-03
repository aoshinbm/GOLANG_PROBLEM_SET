package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	jsonInput := ` {
		"Thinly Sliced Peeled Apples": "6 cups",
		"sugar": "3/4 cup",
		"flour": "2 teaspoons",
		"cinnamon": "1/4 teaspoons",
		"nutmeg": "1/8 teaspoons",
		"lemon juice": "1 teaspoons"
		}`

	// Declared an empty map
	ingredients := make(map[string]string)

	// Unmarshal or Decode the JSON to the interface.
	err := json.Unmarshal([]byte(jsonInput), &ingredients)
	if err != nil {
		fmt.Println("JSON decode error:", err)
		return
	}
	fmt.Println(ingredients)
	fmt.Printf("%T\n", ingredients)

	fmt.Println("Writing to a file ")

	// in case an error is thrown it is received by the err variable and Fatalf method of
	// log prints the error message and stops program execution
	file, err := os.Create("Ingredients.csv")
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	// Defer is used for purposes of cleanup like closing a running file after the file has
	// been written and main function has completed execution
	defer file.Close()

	// len variable captures the length
	// of the string written to the file.
	detailed, err := file.WriteString(ingredients)
	if err != nil {
		log.Fatalf("failed writing to file: %s", err)
	}

	// Name() method returns the name of the
	// file as presented to Create() method.
	fmt.Printf("\nFile Name: %s", file.Name())
	fmt.Printf("\nLength: %d bytes", detailed)
}
