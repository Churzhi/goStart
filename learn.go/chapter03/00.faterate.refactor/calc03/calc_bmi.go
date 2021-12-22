package calc03

func CalcBMI(tall float64, weight float64) (bmi float64) {
	if tall <= 0 {
		panic("身高和体重不能等于零")
	}
	// todo 验证体重的合法性
	return weight / (tall * tall)
}
