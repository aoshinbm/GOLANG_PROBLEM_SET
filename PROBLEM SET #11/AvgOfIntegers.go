package main

import "fmt"

//function that can take any number of ints as arguments.
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	//Within the function, the type of nums is equivalent to []int.
	//We can call len(nums), iterate over it with range
	for _, num := range nums {
		total = num + total
	}
	fmt.Println(total)
}
func main() {
	//Variadic functions can be called in the usual way with individual arguments.
	sum(1, 2)
	sum(11, 21, 31, 41, 51)

	//If you already have multiple args in a slice,
	//apply them to a variadic function using func(slice...)
	nums := []int{1, 2, 3, 4, 8, 9}
	sum(nums...)
}
