package main

import (
	"fmt"
	"sync"
	"time"
)

type Store struct {
	DataCount int
	Max       int
	lock      sync.Mutex
}

type Producer struct {
}

func (Producer) Produce(s *Store) {
	// 1、锁
	s.lock.Lock()
	// 2、查看库存
	// 3、生产
	// 4、释放锁
	defer s.lock.Unlock()
	if s.DataCount == s.Max {
		fmt.Println("生产者看到库存满了，不生产")
		return
	}
	fmt.Println("开始生产+1")
	s.DataCount++
}

type Consumer struct {
}

func (Consumer) Consume(s *Store) {
	// 1、锁
	s.lock.Lock()
	// 2、查看库存
	// 3、消费
	// 4、释放锁
	defer s.lock.Unlock()
	if s.DataCount == 0 {
		fmt.Println("消费者看到库存没了，不消费")
		return
	}
	fmt.Println("开始消费-1")
	s.DataCount--
}

func main() {

	// 初始化仓库
	s := &Store{Max: 10}
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
