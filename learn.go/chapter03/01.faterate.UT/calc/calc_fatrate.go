// Package calc 使用包时候的名字 可以与文件夹名字不一致
package calc

import "fmt"

func CalcFatRate(bmi float64, age int, sex string) (fatRate float64, err error) {

	if bmi <= 0 {
		err = fmt.Errorf("BMI不能为零或是负数")
	}
	if age > 150 {
		err = fmt.Errorf("年龄超过150岁")
	}
	if age <= 0 {
		err = fmt.Errorf("年龄不能为零或负数")
	}
	if sex != "男" && sex != "女" {
		err = fmt.Errorf("性别输入错误")
	}

	sexWeight := 0
	if sex == "男" {
		sexWeight = 1
	} else {
		sexWeight = 0
	}
	fatRate = (1.2*bmi + getAgeWeight(age)*float64(age) - 5.4 - 10.8*float64(sexWeight)) / 100
	return
}

func getAgeWeight(age int) (ageWeight float64) {
	ageWeight = 0.23
	if age >= 30 && age <= 40 {
		ageWeight = 0.22
	}
	return
}
