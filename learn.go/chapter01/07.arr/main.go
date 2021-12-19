package main

import "fmt"

func main() {
	var age int = 35
	fmt.Println(age)

	var ages [5]int = [5]int{24, 25, 35, 37, 48}
	fmt.Println(ages)

	var ages2 = [5]int{35, 32, 43, 55}
	fmt.Println(ages2)
	ages3 := [5]int{2, 233, 44, 56, 22}
	fmt.Println(ages3)
	ages2 = ages3

	ages4 := [...]int{3, 4, 5, 5, 35, 54, 2}
	fmt.Println(ages4)
	// 数组长度不一致时不可相互赋值

	for i := 0; i < len(ages4); i++ {
		fmt.Println(ages4[i])
	}
	for i, val := range ages3 {
		fmt.Println(ages3[i], "====>", i, val)
	}
}
