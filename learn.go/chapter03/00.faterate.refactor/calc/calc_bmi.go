package calc

import gobmi "github.com/armstrongli/go-bmi"

func CalcBMI(weightKG, tallM float64) (bmi float64) {
	bmi, _ = gobmi.BMI(weightKG, tallM)

	return
}
