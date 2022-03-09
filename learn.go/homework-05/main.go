package main

import (
	"fmt"
	"log"
)

func main() {

	filePath := "C:\\Users\\Churzhi\\go\\src\\learn.go\\homework-05\\personalInformation.json"
	input := &inputFromStd{}
	calc := &Calc{}
	rank := &FatRateRank{}
	record := NewRecord(filePath)

	for {
		personInformation := input.GetInput()
		fatRate, err := calc.FatRate(personInformation)
		personInformation.FatRate = fatRate
		if err != nil {
			log.Fatal("计算体脂率失败：", err)
		}
		if err := record.saveInformation(personInformation); err != nil {
			log.Fatal("保存失败", err)
		}
		rank.inputRecord(personInformation.Name, fatRate)
		rankResult, _ := rank.getRankBubbleSort(personInformation.Name)
		fmt.Println("排名结果：", rankResult)

	}

}
