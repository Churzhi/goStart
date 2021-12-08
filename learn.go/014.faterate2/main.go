package main

import (
	"fmt"
)

/*
计算多个人的平均体脂
实现完整的体脂计算器
连续输入三人的姓名、性别、身高、体重、年龄信息
计算每个人的BMI、体脂率
输出：
每个人的姓名、BMI、体脂率、建议
总人数、平均体脂率
*/

func main() {
	for {

		// 定义用户结构体
		type Person struct {
			name      string
			weight    float64
			tall      float64
			age       int
			sexWeight int
			sex       string
			bmi       float64
		}

		var person Person
		users := [3]Person{person}
		var avgBMI float64
		var totalBMI float64
		fmt.Print("请输入一组用户数据（3个用户一组）：\n")
		for i := 0; i < 3; i++ {
			// 用户输入
			fmt.Printf("请输入第%d个用户数据\n", i+1)
			person.name, person.weight, person.tall, person.age, person.sex = getInfoFromInput()

			for {
				if person.sex != "男" && person.sex != "女" {
					fmt.Println("输入错误，请重新输入！")
					fmt.Println("性别(男/女)：")
					fmt.Scanln(&person.sex)
				} else {
					break
				}
			}
			if person.sex == "男" {
				person.sexWeight = 1
			} else {
				person.sexWeight = 0
			}
			person.bmi = calcBMI(person.tall, person.weight)
			users[i] = person
		}

		for i, personVal := range users {

			var fatRate = (1.2*personVal.bmi + 0.23*float64(personVal.age) - 5.4 - 10.8*float64(personVal.sexWeight)) / 100
			var BMIData string
			if personVal.sex == "男" {
				// 判断男性BMI
				if personVal.age >= 18 && personVal.age <= 39 {
					if fatRate <= 0.1 {
						BMIData = "目前是：偏瘦"
					} else if fatRate > 0.1 && fatRate <= 0.16 {
						BMIData = "目前是：标准"
					} else if fatRate > 0.16 && fatRate <= 0.21 {
						BMIData = "目前是：偏重"
					} else if fatRate > 0.21 && fatRate <= 0.26 {
						BMIData = "目前是：肥胖，少吃点"
					} else {
						BMIData = "目前是：非常肥胖!"
					}
				} else if personVal.age >= 40 && personVal.age <= 59 {
					if fatRate <= 0.11 {
						BMIData = "目前是：偏瘦"
					} else if fatRate > 0.11 && fatRate <= 0.17 {
						BMIData = "目前是：标准"
					} else if fatRate > 0.17 && fatRate <= 0.22 {
						BMIData = "目前是：偏重"
					} else if fatRate > 0.22 && fatRate <= 0.27 {
						BMIData = "目前是：肥胖，少吃点"
					} else {
						BMIData = "目前是：非常肥胖!"
					}
				} else if personVal.age >= 60 {
					if fatRate <= 0.13 {
						BMIData = "目前是：偏瘦"
					} else if fatRate > 0.13 && fatRate <= 0.19 {
						BMIData = "目前是：标准"
					} else if fatRate > 0.19 && fatRate <= 0.24 {
						BMIData = "目前是：偏重"
					} else if fatRate > 0.24 && fatRate <= 0.29 {
						BMIData = "目前是：肥胖，少吃点"
					} else {
						BMIData = "目前是：非常肥胖!"
					}
				} else {
					BMIData = "未成年"
				}
			} else if personVal.sex == "女" {
				// 判断女性BMI
				if personVal.age >= 18 && personVal.age <= 39 {
					if fatRate <= 0.2 {
						BMIData = "目前是：偏瘦"
					} else if fatRate > 0.20 && fatRate <= 0.27 {
						BMIData = "目前是：标准"
					} else if fatRate > 0.27 && fatRate <= 0.34 {
						BMIData = "目前是：偏重"
					} else if fatRate > 0.34 && fatRate <= 0.39 {
						BMIData = "目前是：肥胖，少吃点"
					} else {
						BMIData = "目前是：非常肥胖!"
					}
				} else if personVal.age >= 40 && personVal.age <= 59 {
					if fatRate <= 0.21 {
						BMIData = "目前是：偏瘦"
					} else if fatRate > 0.21 && fatRate <= 0.28 {
						BMIData = "目前是：标准"
					} else if fatRate > 0.28 && fatRate <= 0.35 {
						BMIData = "目前是：偏重"
					} else if fatRate > 0.35 && fatRate <= 0.40 {
						BMIData = "目前是：肥胖，少吃点"
					} else {
						BMIData = "目前是：非常肥胖!"
					}
				} else if personVal.age >= 60 {
					if fatRate <= 0.22 {
						BMIData = "目前是：偏瘦"
					} else if fatRate > 0.22 && fatRate <= 0.29 {
						BMIData = "目前是：标准"
					} else if fatRate > 0.29 && fatRate <= 0.36 {
						BMIData = "目前是：偏重"
					} else if fatRate > 0.36 && fatRate <= 0.41 {
						BMIData = "目前是：肥胖，少吃点"
					} else {
						BMIData = "目前是：肥胖，少吃点"
					}
				} else {
					BMIData = "未成年"
				}
			}
			totalBMI += fatRate
			fmt.Printf("第%d位用户 %s 的体脂率是：%f，%s\n", i+1, personVal.name, fatRate, BMIData)

		}
		var userNumbers = len(users)
		avgBMI = totalBMI / float64(userNumbers)
		fmt.Println(userNumbers, "位用户的平均体脂率为：", avgBMI)
		var whetherContinue string
		fmt.Print("是否录入下一组数据n(y/n):")
		fmt.Scanln(&whetherContinue)
		if whetherContinue != "y" {
			break
		}
	}
}
func getInfoFromInput() (name string, weight float64, tall float64, age int, sex string) {
	// 用户输入
	fmt.Print("姓名：")
	fmt.Scanln(&name)

	fmt.Print("体重（kg）：")
	fmt.Scanln(&weight)

	fmt.Print("身高（cm）：")
	fmt.Scanln(&tall)

	fmt.Print("年龄：")
	fmt.Scanln(&age)

	fmt.Print("性别(男/女)：")
	fmt.Scanln(&sex)

	return name, weight, tall, age, sex
}

func calcBMI(tall float64, weight float64) float64 {
	return weight / ((tall * tall) / 10000)
}
