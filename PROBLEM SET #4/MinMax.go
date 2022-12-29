package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	//declaring array with predefined values
	arrayLiteral := "34, 67, 55, 33, 12, 98, 88, 76"
	fmt.Println(arrayLiteral)
	var delimiter = ","
	var parts = strings.Split(arrayLiteral, delimiter) //splitting the the string
	fmt.Println(parts)
	arrayElements, err := strconv.Atoi(arrayLiteral) //string to integer
	fmt.Println("Array Elements :", arrayElements)
	fmt.Println("Error status : ", err)

	var max int = arrayElements[0]
	for i := 0; i < len(arrayElements); i++ {
		for j := i + 1; j < len(arrayElements); j++ {
			if arrayElements[i] > arrayElements[j] {
				max = arrayElements[i]
			}
		}
	}
	fmt.Println("Maximum :", max)

}
