package main

import (
	"fmt"
	"strings"
)

func main() {
	var place string

	fmt.Println("Enter any place: ")
	fmt.Scanln(&place)
	fmt.Println("PLace name : ", place)

	upper := strings.ToUpper(place)
	fmt.Println("PLace name : ", upper)
	//	upper1 := strings.ToUpperSpecial(place)
	//	fmt.Println("PLace name : ", upper1)

	lower := strings.ToLower(place)
	fmt.Println("PLace name : ", lower)
	//	lower1 := strings.ToLowerSpecial(place)
	//	fmt.Println("PLace name : ", lower1)
}
