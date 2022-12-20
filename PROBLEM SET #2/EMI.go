package main

import (
	"fmt"
	"math"
)

func main() {
	var year, n int
	var principalAmt, rateOfInterest, power, r float64

	fmt.Println("Enter the Principal Amount:")
	fmt.Scan(&principalAmt)

	fmt.Println("Enter the Rate of Interest:")
	fmt.Scan(&rateOfInterest)

	fmt.Println("Enter number of Years:")
	fmt.Scan(&year)

	r = rateOfInterest / (12 * 100)
	fmt.Println("Value of r: ", r)

	n = 12 * year
	fmt.Println("Value of n: ", n)

	m := 1 + r
	q := (-n)
	power = math.Pow(m, float64(q))

	payment := (principalAmt * rateOfInterest) / power
	fmt.Println("Amount : ", payment)

}
