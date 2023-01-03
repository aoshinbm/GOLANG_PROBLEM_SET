package main

import "fmt"

func filterFloatNumbers(numbers []float64, myFunc func(float64) bool) []float64 {
	filteredNumbers := make([]float64, 0)
	for _, bl := range numbers {
		if myFunc(bl) {
			filteredNumbers = append(filteredNumbers, bl)
		}
	}
	return filteredNumbers
}

func main() {

	numbers := []float64{5.12, 15.0, 3.39, 8.0, 9.0, 14.0, 6.63}

	fmt.Println("Numbers which have digit after decimal")
	decimalNumbers := filterFloatNumbers(numbers, func(num []float64) bool {
		return numbers == 0
	})
	fmt.Println(decimalNumbers)

	fmt.Println("---------------------------------------")

	fmt.Println("Numbers which have digit after decimal")
	nonDecimalNumbers := nonDecimalNosFilter(numbers, func(number []float64) bool {
		return numbers == 1
	})
	fmt.Println(nonDecimalNumbers)

	fmt.Println("---------------------------------------")

	fmt.Println("Numbers which are divisible ny 3")
	divisiblee := filterFloatNumbers(numbers, func(nums []float64) bool {
		return (numbers % 3) == 0
	})
	fmt.Println(divisiblee)
}
