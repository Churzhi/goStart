package main

import "fmt"

func main() {

	const pi = 3.14159265
	radius1 := 4
	radius2 := 5

	var area1 = 2 * float64(radius1) * pi
	var area2 = 2 * float64(radius2) * pi
	fmt.Printf("area1: %.3f\n", area1)
	fmt.Printf("area2: %.3f\n", area2)

}
