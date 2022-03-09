package main

import (
	"fmt"
	"math"
	"sort"
)

type RankItem struct {
	Name    string
	FatRate float64
}

type FatRateRank struct {
	items []RankItem
}

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
		r.items = append(r.items, RankItem{
			Name:    name,
			FatRate: minFatRate,
		})
	}
}

func (r *FatRateRank) GetRank(name string) (rank int, fatRate float64) {
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

	// for nameItem, frItem := range personFatRate {
	// 	fatRate2PersonMap[frItem] = append(fatRate2PersonMap[frItem], nameItem)
	// 	rankArr = append(rankArr, frItem)
	// }
	// sort.Float64s(rankArr)
	// for i, frItem := range rankArr {
	// 	_names := fatRate2PersonMap[frItem]
	// 	for _, _name := range _names {
	// 		if _name == name {
	// 			rank = i + 1
	// 			fatRate = frItem
	// 			return
	// 		}
	// 	}
	// }
	return
}

func (r *FatRateRank) getRankBubbleSort(name string) (rank int, fatRate float64) {
	fmt.Println(r.items)
	r.bubbleSort()
	frs := map[float64]struct{}{}
	for _, item := range r.items {
		frs[item.FatRate] = struct{}{}
		if item.Name == name {
			fatRate = item.FatRate
		}
	}
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

func (r *FatRateRank) bubbleSort() {

	for i := 0; i < len(r.items); i++ {
		for j := 0; j < len(r.items)-i-1; j++ {
			if r.items[j].FatRate > r.items[j+1].FatRate {
				r.items[j], r.items[j+1] = r.items[j+1], r.items[j]
			}
		}
		fmt.Println("中间状态", r.items)
	}
	fmt.Println("最终状态", r.items)

}
