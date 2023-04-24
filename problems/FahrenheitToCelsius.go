package main

import "fmt"

func main() {
	fmt.Print("Enter a degree (°F): ")
	var input float64
	fmt.Scanf("%f", &input)
	celsius := (input - 32) * 5.0 / 9.0
	fmt.Println(input, "-> °C =", celsius)
}
