package calc00

import "fmt"

func CalcBMI(tall float64, weight float64) (bmi float64, err error) {
	if tall <= 0 {
		return 0, fmt.Errorf("身高和体重不能等于零")
	}
	// todo 验证体重的合法性
	return weight / (tall * tall), nil
}
