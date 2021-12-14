package main

import "fmt"

func main() {
	panicAndRecover()
	//deferGuess()
	//openFile()
	closureMain()
	fmt.Println(calcSum(12, 3, 5, 6, 3, 3))
	fmt.Println(calcSum(12, 3, 5, 6, 3, 3))
	fmt.Println(calcSum(12, 3, 5, 6, 3, 3))

	showUsedTimes()
	//guess(0, 100)

	tall, weight := 1.70, 70.0
	fmt.Println(tall, weight)

	calculatorAdd := func(a, b int) int {
		return a + b
	}
	result := calculatorAdd(1, 2)
	fmt.Println(result)

	fmt.Println("开始计算：")
	fmt.Println("计算结果：", fib(1))

}

func calcBMI() {
	tall, weight := "1.70", 80.0
	fmt.Println(tall, weight)
}

//递归
func fib(n uint) uint {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}
