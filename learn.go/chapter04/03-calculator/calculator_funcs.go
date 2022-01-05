package main

import "fmt"

// Calculator 的成员函数
// 结构体的成员函数执行时，只能通过结构体指针的成员函数进行更改
func (c Calculator) Add() int {
	fmt.Printf("&c @ calculator.funcs = %p\n", &c)
	tempResult := c.left + c.right
	c.result = tempResult
	fmt.Println("调用Add函数，c 的 result = ", c.result)
	return tempResult
}
func (c Calculator) Sub() int {
	return 0
}
func (c Calculator) Multiple() int {
	return 0
}
func (c Calculator) Divide() int {
	return 0
}
func (c Calculator) Reminder() int {
	return c.left % c.right
}

// SetResult 读写分离，只有 SetResult 可以写该结构体的值
func (c *Calculator) SetResult(result int) {
	c.result = result
}
