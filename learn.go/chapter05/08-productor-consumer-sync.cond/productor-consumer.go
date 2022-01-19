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
	pCond     *sync.Cond
	cCond     *sync.Cond
}

type Producer struct {
}

func (Producer) Produce(s *Store) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.DataCount == s.Max {
		fmt.Println("生产者看到库存满了，不生产")
		s.pCond.Wait()
		return
	}
	fmt.Println("开始生产+1")
	s.DataCount++
	s.cCond.Signal()
}

type Consumer struct {
}

func (Consumer) Consume(s *Store) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.DataCount == 0 {
		fmt.Println("消费者看到库存满了，不消费")
		s.cCond.Wait()
		return
	}
	fmt.Println("消费者消费-1")
	s.DataCount--
	s.pCond.Signal()
}

func main() {

	s := &Store{
		Max: 10,
	}
	// 初始化 cond
	//cond := sync.NewCond(&sync.Mutex{})
	s.pCond = sync.NewCond(&s.lock)
	s.cCond = sync.NewCond(&s.lock)

	pCount, cCount := 50, 50
	for i := 0; i < pCount; i++ {
		go func() {
			for {
				time.Sleep(100 * time.Microsecond)
				Producer{}.Produce(s)
			}
		}()
	}
	for i := 0; i < cCount; i++ {
		go func() {
			for {
				time.Sleep(100 * time.Microsecond)
				Consumer{}.Consume(s)
			}
		}()
	}
	time.Sleep(1 * time.Second)
	fmt.Println(s.DataCount)
}