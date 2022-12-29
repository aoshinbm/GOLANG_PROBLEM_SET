package main

import "fmt"

func main() {

	var integerNumbers = []int{}
	fmt.Println("Enter integer numbers :")
	for i := 0; i < len(integerNumbers); i++ {
		fmt.Printf("%d", integerNumbers[i])
	}

}

/*
What you want to use are called Slices.

An array has a fixed size: [n]T is an array of n values of type T.

A slice, on the other hand, is a dynamically-sized, and flexible way of representing an array: []T is a slice with elements of type T.

Slices are more common in the Go world.

package main

import "fmt"

func main() {
    len := 0
    fmt.Print("Enter the number of floats: ")
    fmt.Scanln(&len)
    input := make([]float64, len)
    for i := 0; i < len; i++ {
        fmt.Print("Enter a float: ")
        fmt.Scanf("%f", &input[i])
    }
    fmt.Println(input)
}*/
