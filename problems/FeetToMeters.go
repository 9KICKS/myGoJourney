package main

import "fmt"

func main() {
	fmt.Print("Enter feet (ft): ")
	var input float64
	fmt.Scanf("%f", &input)
	meters := input * 0.3048
	fmt.Println(input, "-> meters (m) = ", meters)
}
