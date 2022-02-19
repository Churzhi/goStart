package main

import (
	"fmt"
	"sync"
	"time"
)

type Store struct {
	init      sync.Once
	store     chan int
	DataCount int
	Max       int
	lock      sync.Mutex
}

type Producer struct {
}

func (s *Store) instrument() {
	s.init.Do(func() {
		s.store = make(chan int, s.Max)
	})
}

func (Producer) Produce(s *Store) {

	fmt.Println("开始生产+1")

}

type Consumer struct {
}

func (Consumer) Consume(s *Store) {

	fmt.Println("开始消费-1", <-s.store)
}

func main() {

	// 初始化仓库
	s := &Store{Max: 10}
	s.instrument()
	pCount, cCount := 50, 50
	for i := 0; i < pCount; i++ {
		// 开启协程
		go func() {
			for {
				time.Sleep(100 * time.Microsecond)
				Producer{}.Produce(s)
			}
		}()
	}
	for i := 0; i < cCount; i++ {
		// 开启协程
		go func() {
			for {
				time.Sleep(100 * time.Microsecond)
				Consumer{}.Consume(s)
			}
		}()
	}
	// 主线程结束时，goroutine 即使没有运行完成也会结束
	time.Sleep(1 * time.Second)
	fmt.Println(s.DataCount)
}
