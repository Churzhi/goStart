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
	items            []Person
	updateFateRateCh chan float64
}

func main() {
	r := &FatRateRank{}
	Register(1000, r)

	r.updateFateRateCh = make(chan float64, 10000)
	base := randFloat64(0, 0.4)
	wg := sync.WaitGroup{}
	wg.Add(5)

	for i := 0; i < 1000; i++ {
		go func(i int) {

			name := "Person:" + strconv.Itoa(i+1)
			_, fatRate := r.getRank(name)
			//randomNum := (rand.Float64() * 0.02) + 0.02
			for {
				r.updateFateRateCh <- fatRate + base*0.2
				r.updateFateRate(name, r.updateFateRateCh)

				rank, newFatRate := r.getRank(name)
				fmt.Println(name, "排名：", rank, "FatRate：", newFatRate)
				//defer wg.Done()
			}
		}(i)
	}

	wg.Wait()
	//time.Sleep(2 * time.Second)

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

// 修改数据
func (r *FatRateRank) updateFateRate(name string, updateFateRateCh chan float64) {

	found := false
	updateFateRate := <-updateFateRateCh
	for i, item := range r.items {
		if item.Name == name {

			if item.FatRate >= updateFateRate {
				item.FatRate = updateFateRate
			}
			r.items[i] = item
			found = true
			break
		}
	}
	if !found {
		r.items = append(r.items, Person{
			Name:    name,
			FatRate: updateFateRate,
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
	for i := 0; i < Number; i++ {
		personFateRateRand := randFloat64(0, 0.4)
		r.inputRecord("Person:"+strconv.Itoa(i+1), personFateRateRand)
	}

	fmt.Println("注册完成")
}

func randFloat64(min, max float64) float64 {

	rand.Seed(time.Now().UnixNano())
	res := min + rand.Float64()*(max-min)

	return res
}
