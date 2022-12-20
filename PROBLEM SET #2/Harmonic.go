package main

import "fmt"

func main() {
	var numberr int
	var sum float64 = 0.0
	fmt.Println("Enter a number: ")
	fmt.Scanln(&numberr)

	for i := 1; i <= numberr; i++ {
		u := 1 / i
		sum = sum + float64(u)
	}
	fmt.Println("Harmoni number sum is : ", sum)
}
