package calc

// v0.0.2

func CalcBMI(tall float64, weight float64) (bmi float64) {
	if tall <= 0 || weight <= 0 {
		panic("身高和体重不能等于零")
	}
	return weight / (tall * tall)
}
