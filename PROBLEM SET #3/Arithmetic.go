package main

import "fmt"

var a, b int

func sum() int {
	return a + b
}
func difference() int {
	return a - b
}
func arithmetic() {
	sum := a + b
	differencee := a - b
	fmt.Println("SUM :", sum)
	fmt.Println("DIFFERENCE :", differencee)
}

func main() {
	fmt.Println("Enter number1 : ")
	fmt.Scanln(&a)
	fmt.Println("Enter number2 : ")
	fmt.Scanln(&b)

	fmt.Println("SUM :", sum())
	fmt.Println("DIFFERENCE :", difference())
	fmt.Println("Sum and Difference of 2 numbers :", arithmetic())

}
