package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num string

	fmt.Println("Enter any number: ")
	fmt.Scanln(&num)

	number, _ := strconv.Atoi(num)
	answer := number * 10
	fmt.Println(answer)
}
