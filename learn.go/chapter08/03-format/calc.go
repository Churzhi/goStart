package main

import (
	"learn.go/pkg/apis"
	"log"

	gobmi "github.com/armstrongli/go-bmi"
)

type Calc struct {
	continental string
}

func (Calc) BMI(person *apis.PersonalInformation) error {
	bmi, err := gobmi.BMI(person.Weight, person.Tall)
	if err != nil {
		log.Println("error when calculating bmi:", err)
		return err
	}
	return nil
}

func (Calc) FatRate(person *apis.PersonalInformation) error {
	person.fatRate = gobmi.CalcFatRate(person.bmi, person.age, person.sex)
	return nil
}
