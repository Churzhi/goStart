package main

import (
	gobmi "github.com/armstrongli/go-bmi"
	"learn.go/pkg/apis"
	"log"
)

type Calc struct {
	continental string
}

// 在 person 这个对象上读写
func (c *Calc) BMI(person *apis.PersonalInformation) (float64, error) {
	bmi, err := gobmi.BMI(person.Weight, person.Tall)
	if err != nil {
		log.Panicln("error when calculating bmi: ", err)
		return -1, err
	}

	return bmi, nil
}

func (c *Calc) FatRate(person *apis.PersonalInformation) (float64, error) {
	bmi, err := c.BMI(person)
	if err != nil {
		return -1, err
	}
	return gobmi.CalcFatRate(bmi, person.Age, person.Sex), err

}
