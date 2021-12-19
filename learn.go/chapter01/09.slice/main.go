package main

import "fmt"

func main() {
	//切片
	a := []int{}
	fmt.Println(a)
	fmt.Println("追加元素到a中，a是切片")
	a = append(a, 2222)
	//数组
	b := [0]int{}
	fmt.Println(b)
	fmt.Println("追加元素到b中，b是数组")
	xzInfo := []string{"xiaozhou", "男"}
	fmt.Println(xzInfo)
	for i, v := range xzInfo {
		fmt.Println(i, v)
	}

	fmt.Println("删除切片中的元素")
	a = []int{111, 222, 333, 444, 555}
	fmt.Println("删除前a：", a)
	a = append(a[:2], a[3:]...)
	fmt.Println("删除后a：", a)

	//插入元素
	backup := append([]int{}, a[1:]...)
	a = append(a[:1], 999)
	a = append(a, backup...)
	fmt.Println(a)
}
