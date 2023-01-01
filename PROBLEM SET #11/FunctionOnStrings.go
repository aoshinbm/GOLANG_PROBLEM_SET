package main

import (
	"fmt"
)

func characterFilter(str []string) []string {

	charcterString := []int{}
	for _, strrings := range str {
		if strrings[i] == "b" {
			charcterString = append(charcterString, strrings)
		}
	}
	return charcterString
}

func threeLettersFilter(str []string) []string {

	threeLetters := []int{}
	for _, strrings := range str {
		if len(strrings) == 3 {
			threeLetters = append(threeLetters, strrings)
		}
	}
	return threeLetters
}

func main() {

	str := []string{"ant", "beetle", "bee", "wasp", "butterfly", "fly", "moth"}
	fmt.Println("---------------------------------------")

	evenNumbers := eventFilter(numbers)
	fmt.Println(evenNumbers)

	fmt.Println("---------------------------------------")

	oddNumbers := oddFilter(numbers)
	fmt.Println(oddNumbers)

}

/* function  -- ()

func functionName (paramtere list -- Any data type(User defiend or golang provide) ) retun type {

}

Higher order function --- () --- Abstracting the logic

func functionName (paramtere list -- another func or func's) ) retun func {

}



*/
