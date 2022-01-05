package main

import (
	gobmi "github.com/armstrongli/go-bmi"
	"log"
)

type Calc struct {
}

// 在 person 这个对象上读写
func (Calc) BMI(person *Person) error {
	bmi, err := gobmi.BMI(person.weight, person.tall)
	if err != nil {
		log.Panicln("error when calculating bmi: ", err)
		return err
	}
	person.bmi = bmi
	return nil
}

func (Calc) FatRate(person *Person) error {
	person.fatRate = gobmi.CalcFatRate(person.bmi, person.age, person.sex)
	return nil
}
