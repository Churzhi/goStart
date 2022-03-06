package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"
)

type Person struct {
	name    string
	fatRate float64
}

type FatRateRank struct {
	personCh chan Person
	items    []Person
}

func main() {
	useChan()
}

func (p *Person) Register(name string, fatRate float64) {
	p.name = name
	p.fatRate = fatRate
}

func useChan() {
	personCount := 1000
	p := Person{}
	r := &FatRateRank{}

	r.personCh = make(chan Person, personCount)
	go func() {
		for i := 0; i < personCount; i++ {
			p.name = "Person" + strconv.Itoa(i)
			// rand.Int 返回一个在[0, max)区间服从均匀分布的随机值，如果max<=0则会panic
			base, err := rand.Int(rand.Reader, big.NewInt(int64(60)))
			if err != nil {
				fmt.Println("rand.Int: error")
			}
			p.fatRate = float64(base.Uint64()) / 100
			r.personCh <- p

			//fmt.Println(p)
		}
	}()

	go func() {
		for range r.personCh {

		}
	}()

}
