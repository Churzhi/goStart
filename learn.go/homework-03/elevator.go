package main

type Elevator struct {
	maxFloor     int   // 楼层总数
	nowFloor     int   // 电梯当前所在楼层 初始默认为1层
	targetFloors []int // 需要 run 的目标楼层
}

func (e *Elevator) OpenDoor() string {
	return "电梯开门"
}

func (e *Elevator) CloseDoor() string {
	return "电梯关门"
}

//func (e *Elevator) ElevatorStatusInfo() string {
//
//	//nowFloorInfo:= "电梯在："+e.nowFloor
//	//
//	return "nil"
//}
