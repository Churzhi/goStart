package calc

import "testing"

//录入正常 BMI、年龄、性别，确保计算结果符合预期；
//录入非法 BMI、年龄、性别（0、负数、超过 150 的年龄、非男女的性别输入），返回错误；
func TestCalcFatRateAge1(t *testing.T) {
	_, err := CalcFatRate(1, -1, "男")
	t.Log("should be error， error is ", err)
	if err == nil {
		t.Errorf("should be error, but err returned is nil")
	}
}
func TestCalcFatRateAge2(t *testing.T) {
	_, err := CalcFatRate(1, 0, "男")
	t.Log("should be error， error is ", err)
	if err == nil {
		t.Errorf("should be error, but err returned is nil")
	}
}
func TestCalcFatRateAge3(t *testing.T) {
	_, err := CalcFatRate(1, 151, "男")
	t.Log("should be error， error is ", err)
	if err == nil {
		t.Errorf("should be error, but err returned is nil")
	}
}
func TestCalcFatRateSex(t *testing.T) {
	_, err := CalcFatRate(1, 25, "1")
	t.Log("should be error， error is ", err)
	if err == nil {
		t.Errorf("should be error, but err returned is nil")
	}
}
func TestCalcFatRateBMI1(t *testing.T) {
	_, err := CalcFatRate(0, 25, "男")
	t.Log("should be error， error is ", err)
	if err == nil {
		t.Errorf("should be error, but err returned is nil")
	}
}
func TestCalcFatRateBMI2(t *testing.T) {
	_, err := CalcFatRate(-1, 25, "男")
	t.Log("should be error， error is ", err)
	if err == nil {
		t.Errorf("should be error, but err returned is nil")
	}
}

func TestCalcFatRateExpected(t *testing.T) {
	var expected float64 = 0.1559
	fatRate, err := CalcFatRate(21.7, 25, "男")
	t.Log("fatRate:", fatRate)
	if err != nil {
		t.Errorf("error is :%v", err)
	}
	if fatRate != expected {
		t.Errorf("expected is different from calculate")
	}
}
