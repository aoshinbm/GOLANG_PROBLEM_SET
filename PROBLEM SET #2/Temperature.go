package main

import "fmt"

func main() {
	var temperature, celsius float64
	var choice int
	fmt.Print("\nEnter the option: \n1.Celsius to Fahrenheit \n2.Fahrenheit to Celsius")
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		fmt.Print("Enter the temperature in Celsius:")
		fmt.Scanln(&temperature)
		f := (float32(temperature) * 1.8) + 32
		fmt.Println("Temperature in Fahrenheit is:", f)
	case 2:
		fmt.Println("Please enter temperature in fahrenheit:")
		fmt.Scanln(&temperature)
		celsius = (temperature - 32) * 5 / 9
		fmt.Println("Temperature in celsius: ", celsius)
	default:
		fmt.Println("INVALID!!!!!!!!!!!")
	}

}
