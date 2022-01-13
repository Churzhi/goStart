package main

import (
	"fmt"
	"time"
)

func ElevatorInit(maxFloor int, nowFloor int) *Elevator {
	ele := new(Elevator)
	ele.maxFloor = maxFloor
	ele.nowFloor = nowFloor

	return ele
}

func (e *Elevator) Run(target int) {
	floor := e.nowFloor
	fmt.Println("电梯当前处于：", floor, "楼,", "正在前往：", target, "楼")
	if target >= 1 && target <= e.maxFloor {
		if target == e.nowFloor {
			fmt.Println("电梯已在你的目标楼层：", target, "楼")
		} else if target > e.nowFloor {
			for elestart := e.nowFloor; elestart < target; elestart++ {
				time.Sleep(1 * time.Second)
				e.nowFloor = elestart + 1
			}
		} else if target < e.nowFloor {
			for elestart := e.nowFloor; elestart > target; elestart-- {
				time.Sleep(1 * time.Second)
				e.nowFloor = elestart - 1
			}
		}
		if e.nowFloor == target {
			fmt.Println(target, "楼到了")
			fmt.Println(e.OpenDoor())
			fmt.Println(e.CloseDoor())
		}
	} else {
		fmt.Println("你的目标楼层错误")
	}
}
