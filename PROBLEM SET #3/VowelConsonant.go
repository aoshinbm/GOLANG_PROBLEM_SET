package main

import (
	"fmt"
	"strings"
)

func main() {
	var alpha string
	var vowel int = 1
	var conson int = 1

	fmt.Println("Enter string: ")
	fmt.Scanln(&alpha)

	split := strings.Split(alpha, ",")
	fmt.Println(split)

	arrayString := [5]string{"a", "e", "i", "o", "u"}
	fmt.Println(arrayString)

	for int i = 0; i < len(split); i++ {
		if split[i] != arrayString[i] {
			conson++
		} else {
			vowel++
		}
	}
	fmt.Println("Vowel : ", vowel)
	fmt.Println("Consonant : ", conson)

}
