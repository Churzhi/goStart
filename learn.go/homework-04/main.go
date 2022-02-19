package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"strconv"
	"sync"
	"time"
)

type Person struct {
	Name    string
	FatRate float64
}

type FatRateRank struct {
	items []Person
}

// 注册数据
func (r *FatRateRank) inputRecord(name string, fatRate ...float64) {
	minFatRate := math.MaxFloat64
	for _, item := range fatRate {
		if minFatRate > item {
			minFatRate = item
		}
	}

	found := false
	for i, item := range r.items {
		if item.Name == name {
			if item.FatRate >= minFatRate {
				item.FatRate = minFatRate
			}
			r.items[i] = item
			found = true
			break
		}
	}
	if !found {
		r.items = append(r.items, Person{
			Name:    name,
			FatRate: minFatRate,
		})
	}
}

// 获取排行榜
func (r *FatRateRank) getRank(name string) (rank int, fatRate float64) {

	sort.Slice(r.items, func(i, j int) bool {
		return r.items[i].FatRate < r.items[j].FatRate
	})
	frs := map[float64]struct{}{}
	for _, item := range r.items {
		frs[item.FatRate] = struct{}{}
		if item.Name == name {
			fatRate = item.FatRate
		}
	}
	rand.Seed(int64(fatRate))
	rankArr := make([]float64, 0, len(frs))
	for k := range frs {
		rankArr = append(rankArr, k)
	}
	sort.Float64s(rankArr)
	for i, frItem := range rankArr {
		if frItem == fatRate {
			rank = i + 1
			break
		}
	}
	return

}

func main() {
	r := &FatRateRank{}
	Register(1000, r)
	rank, _ := r.getRank("Person:22")
	fmt.Println(rank)

	lock := sync.Mutex{}
	concurrency := 1000
	//wg := sync.WaitGroup{}
	//wg.Add(concurrency)

	rand.Seed(time.Now().UnixNano())
	go func() {
		for {
			for _, item := range r.items {
				if item.Name == "" {
					break
				}
				if item.Name == "Person:"+strconv.Itoa(concurrency+1) {
					//defer wg.Done()
					lock.Lock()
					defer lock.Unlock()
					tmpFatRate := item.FatRate
					item.FatRate = tmpFatRate + (rand.Float64() * 0.2) + 0.2
					rank, _ = r.getRank(item.Name)
					fmt.Println(item.Name, "排名：", rank)
				}
			}
		}
	}()

	//wg.Wait()
	time.Sleep(10 * time.Second)

}

func Register(Number int, r *FatRateRank) {
	fmt.Println("开始注册")
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < Number; i++ {
		r.inputRecord("Person:"+strconv.Itoa(i+1), (rand.Float64()*0.5)+0.05)
	}
	//fmt.Println(r)
}
