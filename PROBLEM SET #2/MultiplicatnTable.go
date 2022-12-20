package main

import "fmt"

func main() {
	var number int
	fmt.Println("Enter any number between 2-25:")
	fmt.Scan(&number)

	for i := 0; i <= 10; i++ {
		result := number * i
		fmt.Printf("%d X %d = %d \n", number, i, result)

	}
}
