package main

import "fmt"

func main() {
	var input string
	_, err := fmt.Scanf("%s", &input)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Hello", input)
}
