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

		fmt.Print("性别(男/女)：")
		fmt.Scanln(&sex)
		for {
			if sex != "男" && sex != "女" {
				fmt.Println("输入错误，请重新输入！")
				fmt.Println("性别(男/女)：")
				fmt.Scanln(&sex)
			} else {
				break
			}
		}

		if sex == "男" {
			sexWeight = 1
		} else {
			sexWeight = 0
		}
		fmt.Println(name, weight, tall, age, sex)
		var bmi = weight / (tall * tall)
		var fatRate = (1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*float64(sexWeight)) / 100
		fmt.Printf("体脂率是：%f", fatRate)

		if sex == "男" {
			// 判断男性BMI
			if age >= 18 && age <= 39 {
				if fatRate <= 0.1 {
					fmt.Println("目前是：偏瘦")
				} else if fatRate > 0.1 && fatRate <= 0.16 {
					fmt.Println("目前是：标准")
				} else if fatRate > 0.16 && fatRate <= 0.21 {
					fmt.Println("目前是：偏重")
				} else if fatRate > 0.21 && fatRate <= 0.26 {
					fmt.Println("目前是：肥胖，少吃点")
				} else {
					fmt.Println("目前是：非常肥胖!")
				}
			} else if age >= 40 && age <= 59 {
				if fatRate <= 0.11 {
					fmt.Println("目前是：偏瘦")
				} else if fatRate > 0.11 && fatRate <= 0.17 {
					fmt.Println("目前是：标准")
				} else if fatRate > 0.17 && fatRate <= 0.22 {
					fmt.Println("目前是：偏重")
				} else if fatRate > 0.22 && fatRate <= 0.27 {
					fmt.Println("目前是：肥胖，少吃点")
				} else {
					fmt.Println("目前是：非常肥胖!")
				}
			} else if age >= 60 {
				if fatRate <= 0.13 {
					fmt.Println("目前是：偏瘦")
				} else if fatRate > 0.13 && fatRate <= 0.19 {
					fmt.Println("目前是：标准")
				} else if fatRate > 0.19 && fatRate <= 0.24 {
					fmt.Println("目前是：偏重")
				} else if fatRate > 0.24 && fatRate <= 0.29 {
					fmt.Println("目前是：肥胖，少吃点")
				} else {
					fmt.Println("目前是：非常肥胖!")
				}
			} else {
				fmt.Println("未成年")
			}
		} else if sex == "女" {
			// 判断女性BMI
			if age >= 18 && age <= 39 {
				if fatRate <= 0.2 {
					fmt.Println("目前是：偏瘦")
				} else if fatRate > 0.20 && fatRate <= 0.27 {
					fmt.Println("目前是：标准")
				} else if fatRate > 0.27 && fatRate <= 0.34 {
					fmt.Println("目前是：偏重")
				} else if fatRate > 0.34 && fatRate <= 0.39 {
					fmt.Println("目前是：肥胖，少吃点")
				} else {
					fmt.Println("目前是：非常肥胖!")
				}
			} else if age >= 40 && age <= 59 {
				//todo
				if fatRate <= 0.21 {
					fmt.Println("目前是：偏瘦")
				} else if fatRate > 0.21 && fatRate <= 0.28 {
					fmt.Println("目前是：标准")
				} else if fatRate > 0.28 && fatRate <= 0.35 {
					fmt.Println("目前是：偏重")
				} else if fatRate > 0.35 && fatRate <= 0.40 {
					fmt.Println("目前是：肥胖，少吃点")
				} else {
					fmt.Println("目前是：非常肥胖!")
				}
			} else if age >= 60 {
				//todo
				if fatRate <= 0.22 {
					fmt.Println("目前是：偏瘦")
				} else if fatRate > 0.22 && fatRate <= 0.29 {
					fmt.Println("目前是：标准")
				} else if fatRate > 0.29 && fatRate <= 0.36 {
					fmt.Println("目前是：偏重")
				} else if fatRate > 0.36 && fatRate <= 0.41 {
					fmt.Println("目前是：肥胖，少吃点")
				} else {
					fmt.Println("目前是：非常肥胖!")
				}
			} else {
				fmt.Println("未成年")
			}
		}

		var whetherContinue string
		fmt.Print("是否录入下一个(y/n):")
		fmt.Scanln(&whetherContinue)
		if whetherContinue != "y" {
			break
		}
	}
}
