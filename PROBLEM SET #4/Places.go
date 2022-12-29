package main

import (
	"fmt"
)

func main() {
	var place [7]string
	var count int = 0
	fmt.Println("Enter names of 7 places")
	for i := 0; i < 7; i++ {
		fmt.Scanf("%s", &place[i])
	}
	fmt.Println(place)

	for i := 1; i <= len(place); i++ {
		for j := 1; j <= len(place); j++ {
			count++
			fmt.Println("Count : ", count)
		}
	}

	/*	for _, token := range strings.Split(place, " ") {
			placeSlice = append(placeSlice, place)
		}

	*/
}
