package main

import "fmt"

type Math = int
type English = int

func main() {
	var mathScore int = 100
	var mathScore2 Math = 100
	mathScore2 = mathScore
	fmt.Println(mathScore2)
}
