package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"sync"
	"time"
)

func main() {
	//useChannel()
	fmt.Println("------------------------")
	useLock()
}

func useLock() {
	personCount := 1000
	base := rands(float64(0), 0.4)
	lock := sync.Mutex{}
	persons := Persons{}
	m := sync.Map{}
	wg := sync.WaitGroup{}
	wg.Add(personCount)
	for i := 0; i < personCount; i++ {
		for {
			go func(num int) {
				lock.Lock()
				defer wg.Done()
				defer lock.Unlock()
				name := "person_" + strconv.Itoa(num)
				fatRateUpdate := rands(base-0.2, base+0.2)
				p := &Person{}
				p.register(name, fatRateUpdate)
				persons = append(persons, p)
			}(i)
		}

	}
	wg.Wait()

	//排序
	sort.Sort(persons)
	for index, person := range persons {
		m.Store(person.name, index)
	}
	for i := 0; i < personCount; i++ {
		for {
			go func(num int) {
				name := "person_" + strconv.Itoa(num)
				if rank, ok := m.Load(name); ok {
					fmt.Println(name, "当前排名第：", rank)
				}
			}(i)
		}

	}
	time.Sleep(2 * time.Second)
}

func useChannel() {
	personCount := 1000
	base := rands(float64(0), 0.4)
	chanPerson := make(chan *Person, personCount)
	wg := sync.WaitGroup{}
	wg.Add(personCount)
	for i := 0; i < personCount; i++ {
		go func(num int) {
			defer wg.Done()
			name := "person_" + strconv.Itoa(num)
			fatRateUpdate := rands(base-0.2, base+0.2)
			p := &Person{}
			p.register(name, fatRateUpdate)
			chanPerson <- p
		}(i)
	}
	wg.Wait()
	close(chanPerson)
	maps := updateFatRateLeaderboard(chanPerson)
	for i := 0; i < personCount; i++ {
		go func(num int) {
			name := "person_" + strconv.Itoa(num)
			if rank, ok := maps.Load(name); ok {
				fmt.Println(name, "当前排名第：", rank)
			}
		}(i)
	}
	time.Sleep(2 * time.Second)
}

func updateFatRateLeaderboard(ch chan *Person) (m sync.Map) {
	m = sync.Map{}
	persons := Persons{}
	for person := range ch {
		persons = append(persons, person)
	}
	//排序
	sort.Sort(persons)
	for index, person := range persons {
		m.Store(person.name, index)
	}
	return
}

func rands(min, max float64) (convert float64) {
	max = max - min
	rand.Seed(time.Now().UnixNano())
	convert = roundConvert(min+max*rand.Float64(), 2)
	return
}

func roundConvert(f float64, n int) (res float64) {
	floatStr := fmt.Sprintf("%."+strconv.Itoa(n)+"f", f)
	res, err := strconv.ParseFloat(floatStr, 64)
	if err == nil {
		return
	}
	return
}

type Person struct {
	name    string
	fatRate float64
}

type Persons []*Person

func (p *Person) register(name string, fatRate float64) {
	p.name = name
	p.fatRate = fatRate
}
func (p Persons) Len() int {
	return len(p)

}
func (p Persons) Less(i, j int) bool {
	return p[i].fatRate > p[j].fatRate
}

func (p Persons) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
