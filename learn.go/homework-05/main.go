package main

import (
	gobmi "github.com/armstrongli/go-bmi"
	"log"
)

type Person struct {
	name   string
	sex    string
	age    int
	tall   float64
	weight float64

	bmi     float64
	fatRate float64
}

type Rank interface {
	UpdateFR(name string, fr float64)
	GetRank(name string) int
}

//type Client interface {
//	UpdateFR(name string, fr float64)
//	GetRank(name string) int
//}

type FatRateRank struct {
}

// 实现 Rank 接口方法 UpdateFR()
func (r *FatRateRank) UpdateFR(name string, fr float64) {

}

// 实现 Rank 接口方法 GetRank()
func (r *FatRateRank) GetRank(name string) int {

	return 0
}

func (p *Person) calcBmi() error {
	bmi, err := gobmi.BMI(p.weight, p.tall)
	if err != nil {
		log.Printf("error when calculating BMI for Person[%s]: %v", p.name, err)
		return err
	}
	p.bmi = bmi
	return nil
}

func (p *Person) calcFatRate() {
	p.fatRate = gobmi.CalcFatRate(p.bmi, p.age, p.sex)
}

func main() {

	var rank Rank = &FatRateRank{} // 实例化
	//var clients []Client           // 实例化
	person := getFakePersonInfo()
	person.calcBmi()
	person.calcFatRate()

	rank.UpdateFR(person.name, person.fatRate)

	//for i := 0; i < len(clients); i++ {
	//	go func(idx int) {
	//		// todo add context to control exit
	//		go func() {
	//			for {
	//				clients[idx].UpdateFR("zhou", 0.24)
	//			}
	//		}()
	//
	//		go func() {
	//			for {
	//				clients[idx].GetRank("zhou")
	//			}
	//		}()
	//	}(i)
	//}

}

func getFakePersonInfo() *Person {
	return &Person{
		name:   "zhou",
		sex:    "男",
		tall:   1.74,
		weight: 66,
		age:    24,
	}
}
