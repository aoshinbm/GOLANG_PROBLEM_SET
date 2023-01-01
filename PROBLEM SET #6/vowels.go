package main

import (
	"fmt"
)

func vowel(str string) {
	var count int = 0
	vowels := [5]string{"a", "e", "i", "o", "u"}
	//	fmt.Println("vowels :", vowels)
	for i := 0; i < len(str); i++ {
		if vowels[i] == str {
			count++
			fmt.Printf("%v", count)
		}
	}
}

/*
func main() {

		var sent []string

		vowel(sent)
	}
*/
func main() {
	//	var count int = 0
	var sentence string
	fmt.Println("Enter string : ")
	fmt.Scanln(&sentence)

	vowel(sentence)

	/*	substr := []string{"a", "e", "i", "o", "u"}
		for i := 0; i < len(sentence); i++ {
			check := strings.Contains(sentence, substr[i])
			fmt.Println(check)

		}
	*/
}
