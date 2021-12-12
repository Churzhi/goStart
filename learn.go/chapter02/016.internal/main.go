package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	//arr := []int{}

	// defer 根据定义顺序从后往前执行
	startTime := time.Now()
	defer func() {
		finishTime := time.Now()
		fmt.Println(startTime, finishTime)
		fmt.Println(finishTime.Sub(startTime))
	}()

	arr2 := make([]string, 3, 4)
	fmt.Println("len: ", len(arr2), "cap: ", cap(arr2))
	fmt.Println("default:", arr2[0])
	fmt.Println("default:", arr2[1])
	fmt.Println("default:", arr2[2])

	//i := new()

	arr3 := make([]string, 3, 4)
	//复制切片
	copy(arr3, arr2)
	i := 3333
	j := &i
	fmt.Println(reflect.TypeOf(j))

}
