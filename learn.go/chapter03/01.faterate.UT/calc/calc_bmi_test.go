package calc

// run test 时报错： undefined: CalcBMI；把 vendor目录删除后 run test ok!
import (
	"testing"
)

// 测试录入 0 或负数体重
func TestBMI1(t *testing.T) {
	_, err := CalcBMI(-1, 1)
	if err == nil {
		t.Errorf("should be error, but err returned is nil")
	}
}

// 测试录入 0 或负数身高
func TestBMI2(t *testing.T) {
	_, err := CalcBMI(1, -1)
	if err == nil {
		t.Errorf("should be error, but err returned is nil")
	}
}

// 测试BMI计算是否符合预期
func TestBMI3(t *testing.T) {
	var expected float64 = 1
	bmi, err := CalcBMI(1, 1)
	if err != nil {
		t.Errorf("error is :%v", err)
	}
	if bmi != expected {
		t.Errorf("expected is different from calculate")
	}
}
