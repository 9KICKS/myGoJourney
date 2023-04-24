package main

import "fmt"

func main() {
	fmt.Print("Enter a degree (°C): ")
	var input float64
	fmt.Scanf("%f", &input)
	fahrenheit := (input * 9.0 / 5.0) + 32
	fmt.Println(input, "-> °F =", fahrenheit)
}
