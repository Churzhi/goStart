package main

import "fmt"

func main() {
	//难以长期维护
	xzInfo := [2]string{"xiaozhou", "男"}
	xlInfo := [2]string{"xiaoli", "男"}
	xsInfo := [2]string{"xiaosu", "男"}

	fmt.Println(xzInfo)
	fmt.Println(xlInfo)
	fmt.Println(xsInfo)
	// 支持动态添加
	newPersonInfo := [...][3]string{
		[3]string{"xiaozhou", "男"},
		[3]string{"xiaoli", "男"},
		[3]string{"xiaoli", "男"},
	}
	for _, val := range newPersonInfo {
		fmt.Println(val)
	}

	fmt.Println("降维度输出")
	for d1, d1val := range newPersonInfo {
		for d2, d2val := range d1val {
			fmt.Println(d1, d1val, d2, d2val)
		}
	}
}
