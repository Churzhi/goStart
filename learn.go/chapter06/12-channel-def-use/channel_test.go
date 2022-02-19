package main

import (
	"fmt"
	"testing"
)

func TestDefChannel(t *testing.T) {
	var sampleMap map[string]int = map[string]int{}
	fmt.Println("sampleMap", sampleMap)
	var intCh chan int // =make(chan int)
	fmt.Println("intCh", intCh)
	intCh = make(chan int, 1)
	fmt.Println("intCh", intCh)

	// channel 操作
	fmt.Println("装入数据")
	intCh <- 3
	fmt.Println("取出数据")
	out := <-intCh
	fmt.Println("数据是：", out)

}

func TestChanPutGet(t *testing.T) {

	intCh := make(chan int) // 创建一个不带 size 的 channel（不带 buffer）
	workerCount := 10
	for i := 0; i < workerCount; i++ {
		go func(i int) {
			intCh <- i
		}(i)
	}

	for o := 0; o < workerCount; o++ {
		go func(o int) {
			out := <-intCh
			fmt.Printf("出%d拿到%d", 0, out)
		}(o)
	}
}

func TestRangeClosedChannel(t *testing.T) {
	intCh := make(chan int, 10)
	intCh <- 2
	intCh <- 3
	intCh <- 4
	intCh <- 5
	intCh <- 6

	close(intCh)
	for o := range intCh {
		fmt.Println("range 取数", o)
	}
}

func TestSelectChannel(t *testing.T) {

}
