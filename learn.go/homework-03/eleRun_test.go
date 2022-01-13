package main

import (
	"testing"
)

//**案例一**
//楼层有 5 层，所有电梯楼层没有人请求电梯，电梯不动
func Testcase1(t *testing.T) {

	// 初始化电梯对象
	ele := ElevatorInit()

	t.Log("Testcase1:", "电梯当前楼层为：", ele.nowFloor)
	if ele.nowFloor != 5 {
		t.Fail()
	}
	if ele.nowFloor != 1 {
		t.Fail()
	}
}

func ElevatorInit() interface{} {

}
