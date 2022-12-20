package main

import "fmt"

func main() {
	var x, y, z int
	fmt.Println("Enter three side of triangle: ")
	fmt.Scanln(&x)
	fmt.Scanln(&y)
	fmt.Scanln(&z)

	if x == y && y == z && x == z {
		fmt.Println("Its an ISOSCELES TRIANGLE")
	} else if x == y && y == z {
		fmt.Println("Its an EQUILATERAL TRIANGLE")
	} else {
		fmt.Println("Its an SCALENE TRIANGLE")
	}
}
