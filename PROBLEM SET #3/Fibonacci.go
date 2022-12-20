package main

import "fmt"

func main() {
	var numm,next int
	var first=1 int64 
	var second=1 int 

	fmt.Println("Fibonacci Problem")
	fmt.Println("Enter any number: ")
	fmt.Scanln(&numm)

	fmt.Println(first)
	fmt.Println(second)
	for i := 0; i < numm; i++ {
		next =first+second
		first=second
		second=next
		fmt.Println(next)
	}
}