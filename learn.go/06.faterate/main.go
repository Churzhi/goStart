package main

import (
	"fmt"
)

func main() {
	for {
		var name string
		var weight float64
		var tall float64
		var age int
		var sexWeight int
		var sex string

		// 用户输入
		fmt.Print("姓名：")

		fmt.Scanln(&name)

		fmt.Print("体重（kg）：")

		fmt.Scanln(&weight)
		fmt.Print("身高（m）：")
		fmt.Scanln(&tall)

		fmt.Print("年龄：")
		fmt.Scanln(&age)

		fmt.Print("性别（男/女）：")
		fmt.Scanln(&sex)
		var bmi = weight / (tall * tall)

		if sex == "男" {
			sexWeight = 1
		} else {
			sexWeight = 0
		}
		fmt.Println(name, weight, tall, age, sex)
		var fatRate = (1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*float64(sexWeight)) / 100
		fmt.Printf("体脂率是：%f", fatRate)

		if sex == "男" {
			// todo 编写男性的体脂率与体脂状态表
			if age >= 18 && age <= 39 {
				if fatRate < 0.1 {
					fmt.Println("目前是：偏瘦")
				} else if fatRate > 0.1 && fatRate <= 0.16 {
					fmt.Println("目前是：标准")
				} else if fatRate > 0.16 && fatRate <= 0.21 {
					fmt.Println("目前是：偏重")
				} else if fatRate >= 0.21 && fatRate <= 0.26 {
					fmt.Println("目前是：肥胖，少吃点")
				} else {
					fmt.Println("目前是：非常肥胖，健身游泳跑步，即刻开始")
				}

			} else if age >= 40 && age <= 59 {
				//todo
			} else if age >= 60 {
				//todo
			} else {
				fmt.Println("未成年")
			}
		} else {
			// todo 编写女性的体脂率与体脂状态表

		}

		var whetherContinue string
		fmt.Print("是否录入下一个(y/n):")
		fmt.Scanln(&whetherContinue)
		if whetherContinue != "y" {
			break
		}
	}
}
