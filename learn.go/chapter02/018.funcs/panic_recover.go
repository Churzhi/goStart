package main

import "fmt"

func panicAndRecover() {

	defer coverPanicUpgraded() // 这样可以抓住严重错误

	//捕获
	defer func() { //抓不住严重错误 panic
		coverPanicUpgraded()
	}()

	defer coverPanic()
	var nameScore map[string]int = nil

	nameScore["zhou"] = 100
}

func coverPanic() { // 未能抓住 panic
	func() {
		if r := recover(); r != nil {
			fmt.Println("系统出严重故障：", r)
		}
	}()
}
func coverPanicUpgraded() { // 能抓住 panic

	//todo 确认代码
	func() {
		if r := recover(); r != nil {
			fmt.Println("系统出严重故障：", r)
		}
	}()
}
