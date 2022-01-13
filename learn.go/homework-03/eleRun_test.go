package main

import "testing"

//**案例一**
//楼层有 5 层，所有电梯楼层没有人请求电梯，电梯不动
func Test1(t *testing.T) {
	// 初始化电梯对象
	ele := ElevatorInit(5, 1)

	t.Log("Testcase1:", "电梯当前楼层为：", ele.nowFloor)
	if ele.maxFloor != 5 {
		t.Fail()
	}
	if ele.nowFloor != 1 {
		t.Fail()
	}
}

//**案例二**
//楼层有 5 层，电梯在 1 层。三楼按电梯。电梯向三楼行进，并停在三楼。
func Test2(t *testing.T) {

	// 初始化电梯对象
	ele := ElevatorInit(5, 1)
	t.Log("Testcase2:", "电梯当前楼层为：", ele.nowFloor)
	if ele.maxFloor != 5 {
		t.Fail()
	}
	ele.Run(3)
	t.Log("Testcase2:", "电梯当前楼层为：", ele.nowFloor)
	if ele.nowFloor != 3 {
		t.Fail()
	}
}

//**案例三**
//楼层有 5 层，电梯在 3 层。上来一些人后，目标楼层： 4 楼、2 楼。电梯先向上到 4 楼，然后转头到 2 楼，最后停在 2 楼。
func Test3(t *testing.T) {

	// 初始化电梯对象
	ele := ElevatorInit(5, 3)
	if ele.maxFloor != 5 {
		t.Fail()
	}
	t.Log("Testcase3:", "电梯当前楼层为：", ele.nowFloor)
	if ele.nowFloor != 3 {
		t.Fail()
	}
	ele.Run(4)
	t.Log("Testcase3:", "电梯当前楼层为：", ele.nowFloor)
	if ele.nowFloor != 4 {
		t.Fail()
	}
	ele.Run(2)
	t.Log("Testcase3:", "电梯当前楼层为：", ele.nowFloor)
	if ele.nowFloor != 2 {
		t.Fail()
	}

}

//**案例四**
//楼层有 5 层，电梯在 3 层。上来一些人后，目标楼层： 4 楼、5 楼、2 楼。电梯先向上到 4 楼，然后到 5 楼，之后转头到 2 楼，最后停在 2 楼。
func Test4(t *testing.T) {

	// 初始化电梯对象
	ele := ElevatorInit(5, 3)
	t.Log("Testcase4:", "电梯当前楼层为：", ele.nowFloor)
	if ele.nowFloor != 3 {
		t.Fail()
	}
	ele.Run(4)
	t.Log("Testcase4:", "电梯当前楼层为：", ele.nowFloor)
	if ele.nowFloor != 4 {
		t.Fail()
	}
	ele.Run(5)
	t.Log("Testcase4:", "电梯当前楼层为：", ele.nowFloor)
	if ele.nowFloor != 5 {
		t.Fail()
	}
	ele.Run(2)
	t.Log("Testcase4:", "电梯当前楼层为：", ele.nowFloor)
	if ele.nowFloor != 2 {
		t.Fail()
	}

}
