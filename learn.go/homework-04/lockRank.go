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
	lock  sync.Mutex
	items []Person
}

// 注册/修改数据
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
			r.lock.Lock()
			if item.FatRate >= minFatRate {
				item.FatRate = minFatRate
			}
			r.items[i] = item
			found = true
			defer r.lock.Unlock()
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

func Register(Number int, r *FatRateRank) {
	fmt.Println("开始注册")
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < Number; i++ {
		r.inputRecord("Person:"+strconv.Itoa(i+1), (rand.Float64()*0.5)+0.05)
	}
	//fmt.Println(r)
}

func main() {
	r := &FatRateRank{}
	Register(1000, r)
	//rank, _ := r.getRank("Person:22")
	//fmt.Println(rank)

	//lock := sync.Mutex{}
	//wg := sync.WaitGroup{}
	//wg.Add(5)
	go func() {
		for {
			rand.Seed(time.Now().UnixNano())
			for _, item := range r.items {
				randomNum := (rand.Float64() * 0.02) + 0.02
				r.inputRecord(item.Name, item.FatRate+randomNum)
				rank, fatRate := r.getRank(item.Name)
				fmt.Println(item.Name, "排名：", rank, "FatRate：", fatRate)
			}
			//defer wg.Done()

		}
	}()

	//wg.Wait()
	time.Sleep(10 * time.Second)

}
