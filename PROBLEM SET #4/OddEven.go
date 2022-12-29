package main

import "fmt"

// Function in which slice
// is passed by value
func oddCheck(num []int) {
	var oddSlice = []int{}
	for i := 0; i < len(num); i++ {
		if num[i]%2 != 0 {
			// Here we only modify the slice
			// Using append function
			// Here, this function only modifies
			// the copy of the slice present in
			// the function not the original slice
			oddSlice = append(oddSlice, num[i])
		}
	}
	fmt.Println("Odd Number Slice : ", oddSlice)
}

// Function in which slice
// is passed by value
func evenCheck(n []int) {
	var evenSlice = []int{}
	for i := 0; i < len(n); i++ {
		if n[i]%2 == 0 {
			// Here we only modify the slice
			// Using append function
			// Here, this function only modifies
			// the copy of the slice present in
			// the function not the original slice
			evenSlice = append(evenSlice, n[i])
		}
	}
	fmt.Println("EVEN Number Slice : ", evenSlice)
}

func main() {
	listtOfNumbers := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	fmt.Println(listtOfNumbers)

	slice10Numbers := listtOfNumbers[5:15]
	fmt.Println(slice10Numbers)

	oddCheck(slice10Numbers)
	evenCheck(slice10Numbers)
}
