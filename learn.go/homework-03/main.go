package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

}

func GetInputTargetFloor() *Elevator {

	var maxFloor int
	fmt.Println("总楼层是：")
	fmt.Scanln(&maxFloor)

	var nowFloor int
	fmt.Println("当前电梯所在楼层是：")

	fmt.Scanln(&nowFloor)

	var targetFloors []int

	// 获取带空格字符串的输入
	input := bufio.NewReader(os.Stdin)
	fmt.Println("请输入需去往的楼层，使用空格分隔:")
	target, err := input.ReadString('\n')

	if err == nil {
		targets := strings.Fields(target) //分割字符返回字符串数组
		for _, item := range targets {
			i, _ := strconv.Atoi(item) // 字符转数字
			targetFloors = append(targetFloors, i)
		}
	}

	return &Elevator{
		maxFloor:     maxFloor,
		nowFloor:     nowFloor,
		targetFloors: targetFloors,
	}

}
