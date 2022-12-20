package main

import "fmt"

func main() {
	var temperature int
	fmt.Print("Enter the temperature in Celsius:")
	fmt.Scanf("%d", &temperature)
	f := (float32(temperature) * 1.8) + 32
	fmt.Println("Temperature in Fahrenheit is:", f)
}
