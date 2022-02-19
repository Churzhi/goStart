package main

import (
	"fmt"
	"sort"
)

type Button struct {
	Floor int
}

type Elevator struct {
	buttons  Buttons
	position int
}

// 类型重定义
type Buttons []*Button

func (b Buttons) Len() int {
	// 获取数据集合元素个数
	return len(b)
}

func (b Buttons) Less(i, j int) bool {
	// 如果 i 索引的数据小于 j 索引的数据，返回 true，且不会调用下面的 Swap()，即数据升序排序。
	return b[i].Floor < b[j].Floor
}

func (b Buttons) Swap(i, j int) {
	// 交换 i 和 j 索引的两个元素的位置
	b[i], b[j] = b[j], b[i]
}

func main() {

	ev := &Elevator{
		buttons: []*Button{
			{Floor: 3},
			{Floor: 1},
			{Floor: 5},
			{Floor: 2},
		},
		position: 2,
	}
	// 对 ev.buttons 进行排序，需要先实现这个接口的方法 Len()、Less(i, j int)、Swap(i, j int)
	sort.Sort(sort.Reverse(ev.buttons))

	fmt.Printf("%+v\n", ev.buttons)

	for _, item := range ev.buttons {
		fmt.Println(item.Floor)
	}
}
