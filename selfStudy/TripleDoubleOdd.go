package main

import "fmt"

func main() {
	x := 19 * 3
	y := 50 * 2
	findTheOdd := y % x
	result := findTheOdd % 3
	fmt.Println(result)
}
