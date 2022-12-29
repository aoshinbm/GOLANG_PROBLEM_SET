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
	var parts = strings.Split(arrayLiteral, delimiter)
	fmt.Println(parts)
	a, _ := strconv.Atoi(arrayLiteral)
	fmt.Println(a)

	fmt.Println("---------------------------")
	var arrayOfNumbers string
	fmt.Println("Enter numbers with comma: ")
	fmt.Scanln(&arrayOfNumbers)

	numbers := strings.Split(arrayOfNumbers, ",")
	fmt.Println(numbers)

}
