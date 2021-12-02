package main

import "fmt"

func main() {
	var fruit = "6 apples"
	var watermelon = false
	if watermelon {
		fruit = "1 apple"
	} else {
		fruit = "7 apples"
	}
	fmt.Println("buy: ", fruit)
}
