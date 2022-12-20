package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var randomNumber [5]int
	var average int
	sum := 0
	min := 10
	max := 50
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		randomNumber[i] = rand.Intn(max-min+1) + min
		fmt.Println(randomNumber[i])
		sum = sum + randomNumber[i]
	}
	fmt.Println("Sum : ", sum)
	average = sum / 5
	fmt.Println("Average : ", average)
}
