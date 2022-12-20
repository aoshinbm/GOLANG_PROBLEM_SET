package main

import "fmt"

func main() {
	var first string
	var second string

	fmt.Println("Enter Your First Name: ")
	// Taking input from user
	fmt.Scanln(&first)

	fmt.Println("Enter Second Last Name: ")
	fmt.Scanln(&second)

	fmt.Print("Full Name is: ")
	fmt.Print(first)
	fmt.Print(second)

	fmt.Print("\n")
	fmt.Println("Full Name is: ", first, " ", second)

	fmt.Printf("Full Name is: %v %v", first, second)
}
