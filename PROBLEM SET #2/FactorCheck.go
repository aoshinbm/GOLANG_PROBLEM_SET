package main

import "fmt"

func main() {
	var input int
	var factor []int
	fmt.Println("Enter a number: ")
	fmt.Scanln(&input)

	for i := 0; i < input; i++ {
		if input%i == 0 {
			factor = i
			fmt.Println(factor)
		}
	}
		fmt.Println(factor)

	//for i := 0; i < input; i++ {
		if factor == 3 {
			fmt.Println("PLING")
		} else if factor == 5 {
			fmt.Println("PLANG")
		} else if factor == 7 {
			fmt.Println("PLONG")
		} else {
			fmt.Printf("%d factors are : %d ", input, factor)
		}
	}
}
