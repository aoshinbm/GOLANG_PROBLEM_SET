package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var coinflip int
	fmt.Println("Coin Flipping")

	rand.Seed(time.Now().UnixNano())
	coinflip = rand.Intn(2)
	if coinflip == 0 {
		fmt.Println("TAILS...")
	} else {
		fmt.Println("HEADS..!!!")
	}
}
