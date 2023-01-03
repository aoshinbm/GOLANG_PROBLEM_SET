package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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
	fmt.Println(&ingredients)

	//convert into json
	jsonBytes, err := json.MarshalIndent(ingredients, "", "   ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonBytes))

	fmt.Println("Writing to a file ")

	//Writing struct type to a JSON file
	err = ioutil.WriteFile("AppleIngredients.json", jsonBytes, 0644)
	if err != nil {
		log.Fatal(err)
	}

	//Reading into struct type from a JSON file
	jsonBytes, err = ioutil.ReadFile("AppleIngredients.json")
	if err != nil {
		log.Fatal(err)
	}
	//
	err = json.Unmarshal(content, &user2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Id:%d, Name:%s, Password:%s, LoggedAt:%v", user2.Id, user2.Name, user2.Password, user2.LoggedAt)
}
