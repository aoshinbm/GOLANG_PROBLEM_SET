package main

import "fmt"

func main() {
	var n int
	var fact int = 1

	fmt.Println("Factorial Problem")
	fmt.Println("Enter any number: ")
	fmt.Scanln(&n)

	for i := 1; i <= n; i++ {
		fact = fact * i
		fmt.Println(fact)
	}
}
