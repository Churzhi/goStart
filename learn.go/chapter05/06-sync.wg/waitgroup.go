package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Runner struct {
	Name string
}

func (r Runner) Run(startPointWg, wg *sync.WaitGroup) {
	fmt.Println(r.Name, "开始跑")
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Duration(rand.Uint64()%10) * time.Second)
	fmt.Println(r.Name, "跑到终点")
}

func main() {
	runnerCount := 10
	runners := []Runner{}
	wg := sync.WaitGroup{}

	wg.Add(runnerCount)

	for i := 0; i < runnerCount; i++ {
		runners = append(runners, Runner{
			Name: fmt.Sprintf("%d", i),
		})
	}
	for _, runnerItem := range runners {
		go runnerItem.Run(&&wg)
	}
	fmt.Println("各就位")
	time.Sleep(1*time.Second)
	fmt.Println("预备跑")


}
