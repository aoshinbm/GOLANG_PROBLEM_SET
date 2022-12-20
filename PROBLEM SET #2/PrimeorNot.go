package main

import "fmt"

func main() {
	var input, factor int
	fmt.Println("Enter a number: ")
	fmt.Scanln(&input)

	for i := 0; i <= input; i++ {
		if input/i != 0 {
			prime := i
			fmt.Println(prime)
		}
	}
}
