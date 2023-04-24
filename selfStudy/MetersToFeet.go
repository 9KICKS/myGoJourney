package main

import "fmt"

func main() {
	fmt.Print("Enter meter (m): ")
	var input float64
	fmt.Scanf("%f", &input)
	feet := input / 0.3048
	fmt.Println(input, "-> feet (ft) =", feet)
}
