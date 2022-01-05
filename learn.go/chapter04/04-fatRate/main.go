package main

import (
	"fmt"
)

func main() {

	person := getFakePersonInfo()
	c := Calc{}
	c.BMI(person)
	c.FatRate(person)
	fmt.Println(person)
	sugg := fatRateSuggestion{}
	suggestion := sugg.GetSuggestion(person)
	fmt.Println(suggestion)
}

func getPersonInfoFromInput() *Person {
	// 录入各项
	var name string
	fmt.Print("姓名：")
	fmt.Scanln(&name)

	var weight float64
	fmt.Print("体重（千克）：")
	fmt.Scanln(&weight)

	var tall float64
	fmt.Print("身高（米）：")
	fmt.Scanln(&tall)
	var age int
	fmt.Print("年龄：")
	fmt.Scanln(&age)

	sex := "男"
	fmt.Print("性别（男/女）：")
	fmt.Scanln(&sex)

	return &Person{
		name:   name,
		sex:    sex,
		tall:   tall,
		weight: weight,
		age:    age,
	}
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
