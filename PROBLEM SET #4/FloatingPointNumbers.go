package main

import "fmt"

func main() {

	var fl_numbers float64
	var enter string
	fmt.Println("Floating point numbers....")
	fmt.Println("Enter floating point numbers: ")
	fmt.Scanf("%v", &enter)
	if enter == "q" || enter == "Q" {
		fmt.Println("Entered number is not float so exit the loop")
		break
	} else {
		fmt.Println("Enter floating point numbers: ")
		fmt.Scanf("%.2f", &fl_numbers)
	}
}
