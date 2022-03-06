package main

import "testing"

type Rank interface {
	UpdateFR(name string, fr float64)
	GetRank(name string) int
}

type Client interface {
	UpdateFR(name string, fr float64)
	GetRank(name string) int
}

func TestHomework(t *testing.T) {
	var rank Rank        // 实例化
	var clients []Client // 实例化

	for i := 0; i < len(clients); i++ {
		go func(idx int) {
			// todo add context to control exit
			go func() {
				for {
					clients[idx].UpdateFR("zhou", 0.24)
				}
			}()

			go func() {
				for {
					clients[idx].GetRank("zhou")
				}
			}()
		}(i)
	}
}
