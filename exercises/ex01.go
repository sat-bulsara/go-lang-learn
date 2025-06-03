package main

import "fmt"

func main() {
	// Declare and initialize temperature and humidity
	var temperature int
	var humidity int

	fmt.Println("What is the temperature in degrees Celsius?")
	fmt.Scanln(&temperature)
	fmt.Println("What is the humidity percentage?")
	fmt.Scanln(&humidity)

	// Check conditions
	if temperature < 10 {
		if humidity > 80 {
			fmt.Println("Wear a raincoat.")
		} else {
			fmt.Println("Wear a warm coat.")
		}
	} else if temperature >= 10 && temperature <= 20 {
		if humidity > 80 {
			fmt.Println("Wear a raincoat.")
		} else {
			fmt.Println("Wear a  coat.")
		}
	} else {
		if humidity > 80 {
			fmt.Println("carry a umbrella.")
		} else {
			fmt.Println("Wear light clothing")
		}
	}
}
