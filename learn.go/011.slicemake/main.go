package main

import "fmt"

func main() {

	{
		a := make([]int, 0)
		fmt.Println("len:", len(a))
		a = append(a, 1)
		fmt.Println("len:", len(a), "val", a)
	}
	{
		//make 1 已经给定了一个长度
		a := make([]int, 1)
		fmt.Println("len:", len(a))
		a = append(a, 1)
		fmt.Println("len:", len(a), "val", a)
	}
	{
		//如果追求性能，建议写法
		a := make([]int, 0, 1)
		fmt.Println("len:", len(a), cap(a))
		a = append(a, 1)
		fmt.Println("len:", len(a), "val", a)
	}
}
