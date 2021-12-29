package calc

import "fmt"

func CalcBMI(weightKG, tallM float64) (bmi float64, err error) {
	if weightKG < 0 {
		err = fmt.Errorf("身高不能玩为负数")
	}
	if tallM < 0 {
		err = fmt.Errorf("体重不能玩为负数")
	}
	if weightKG == 0 {
		err = fmt.Errorf("身高不能玩为零")
	}
	if tallM == 0 {
		err = fmt.Errorf("体重不能玩为零")
	}
	bmi = weightKG / (tallM * tallM)
	return
}
