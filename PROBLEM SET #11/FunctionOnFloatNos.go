package main

import "fmt"

func decimalNosFilter(numbers []float64) []float64 {

	decimalNumbers := []float64{}
	for _, number := range numbers {
		if (number%2) == 0.0 {
			decimalNumbers = append(decimalNumbers, number)
		}
	}
	return decimalNumbers
}

func nonDecimalNosFilter(numbers []float64) []float64 {

	nonDecimalNumbers := []float64{}
	for _, number := range numbers {
		if number != .0 {
			nonDecimalNumbers = append(nonDecimalNumbers, number)
		}
	}
	return nonDecimalNumbers
}

func divisibleby3(num []float64) []float64 {

	divisible := []float64{}
	for _, number := range num {
		if (number % 3) == 0.0 {
			divisible = append(divisible, number)
		}
	}
	return divisible
}

func main() {

	numbers := []float64{5.12, 15.0, 3.39, 8.0, 9.0, 14.0, 6.63}

	fmt.Println("Numbers which have digit after decimal")
	decimalNumbers := decimalNosFilter(numbers)
	fmt.Println(decimalNumbers)

	fmt.Println("---------------------------------------")

	fmt.Println("Numbers which have digit after decimal")
	nonDecimalNumbers := nonDecimalNosFilter(numbers)
	fmt.Println(nonDecimalNumbers)

	fmt.Println("---------------------------------------")

	fmt.Println("Numbers which are divisible ny 3")
	divisiblee := divisibleby3(numbers)
	fmt.Println(divisiblee)
}

/* function  -- ()

func functionName (paramtere list -- Any data type(User defiend or golang provide) ) retun type {

}

Higher order function --- () --- Abstracting the logic

func functionName (paramtere list -- another func or func's) ) retun func {

}
*/
