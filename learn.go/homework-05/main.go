package main

import (
	"encoding/json"
	"fmt"
	gobmi "github.com/armstrongli/go-bmi"
	"io/ioutil"
	"log"
	"os"
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

	//var rank Rank = &FatRateRank{} // 实例化
	//var clients []Client           // 实例化s
	input := &inputFromStd{}
	personInformation := input.GetInput()
	filePath := "C:\\Users\\Churzhi\\go\\src\\learn.go\\homework-05\\personInformation.json"

	data, err := json.Marshal(personInformation)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println("Marshal 的结果是(原生)：", data)
	fmt.Println("Marshal 的结果是(string)：", string(data))

	//保存注册信息到文件中
	writeFile(filePath, data)
	readFile(filePath)

	//rank.UpdateFR(person.name, person.fatRate)

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

func writeFile(filePath string, data []byte) {
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("无法打开文件", filePath, "错误信息是：", err)
		os.Exit(1)
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func readFile(filePath string) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("读取文件失败：", err)
		return
	}
	fmt.Println("读取出来的内容是：", string(data))

}
