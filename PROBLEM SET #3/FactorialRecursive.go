package main

import "fmt"

var n int
var fact int = 1

func factorial() {
	for i := 1; i <= n; i++ {
		fact = fact * i
		fmt.Println(fact)
	}
}

func main() {
	fmt.Println("Factorial Problem")
	fmt.Println("Enter any number: ")
	fmt.Scanln(&n)

	factorial()

}
