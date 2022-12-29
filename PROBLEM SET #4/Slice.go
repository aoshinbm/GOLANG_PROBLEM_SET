package main

import (
	"fmt"
	"math"
)

func power(inpSlice []uint8) {
	powerOf2Slice := []uint8{}
	for i := 0; i < len(inpSlice); i++ {
		value := math.Pow(float64(inpSlice[i]), 2)
		a := uint8(value)
		powerOf2Slice = append(powerOf2Slice, a)
	}
	fmt.Println(inpSlice, "==>", powerOf2Slice)
}
func main() {
	inpSlice := make([]uint8, 3, 5)
	fmt.Println(inpSlice)

	inpSlice[0] = 2
	inpSlice[1] = 4
	inpSlice[2] = 8
	fmt.Println(inpSlice)

	power(inpSlice)

}
