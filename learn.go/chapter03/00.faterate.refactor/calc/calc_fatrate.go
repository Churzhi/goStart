// Package calc 使用包时候的名字 可以与文件夹名字不一致
package calc

import gobmi "learn.go/staging/src/github.com/armstrongli/go-bmi"

func CalcFatRate(bmi float64, age int, sex string) (fatRate float64) {
	return gobmi.CalcFatRate(bmi, age, sex)
}
