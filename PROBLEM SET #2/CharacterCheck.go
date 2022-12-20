package main

import (
	"fmt"
	"strings"
)

func main() {
	var user string
	fmt.Println("Enter a string: ")
	fmt.Scanln(&user)

	/*
		strings.Contains Function in Golang is used to check
		the given letters present in the given string or not.
		If the letter is present in the given string,
		then it will return true, otherwise, return false.
	*/

	if strings.Contains(user, "a") && strings.Contains(user, "e") && strings.Contains(user, "p") {
		fmt.Println("All Present....")
	} else if strings.Contains(user, "a") || strings.Contains(user, "e") || strings.Contains(user, "p") {
		fmt.Println("One or More - Present....")
	} else {
		fmt.Println("None Present")
	}

}
