package calc

func CalcBMI(weightKG, tallM float64) (bmi float64) {

	bmi = weightKG / (tallM * tallM)
	return
}
