package main

type Elevator struct {
	maxFloor     int          // 楼层总数
	nowFloor     int          // 电梯当前所在楼层 初始默认为1层
	targetFloors map[int]bool // 需要 run 的目标楼层
	direction    string       // 方向
}

func (e *Elevator) OpenDoor() string {
	return "电梯开门"
}

func (e *Elevator) CloseDoor() string {
	return "电梯关门"
}
