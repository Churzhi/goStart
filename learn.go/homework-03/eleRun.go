package main

import (
	"fmt"
	"time"
)

type eleTask struct {
	e *Elevator
}

func (eT eleTask) Run(ele Elevator) {
	for _, item := range ele.targetFloors {

		if ele.nowFloor < item {
			eT.passingFloor(ele.nowFloor, item)
		}
		if ele.nowFloor == item {
			eT.reachTargetFloor(ele.nowFloor)
		}
	}

}

func (eT *eleTask) passingFloor(passing, target int) {
	fmt.Println("当前楼层：", passing, "楼", "去往", target, "楼")
	time.Sleep(1000)
}

func (eT *eleTask) reachTargetFloor(target int) {
	fmt.Println("到达：", target, "楼")
	eT.e.OpenDoor()
	time.Sleep(1000)
	eT.e.OpenDoor()
}
