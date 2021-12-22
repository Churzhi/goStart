package calc00

import "testing"

func TestCalcBMI(t *testing.T) {
	inputHeight, inputWeight := 1.0, 1.0
	expectedOutput := 1.0
	t.Logf("开始计算，输入：height：%f,weight: %f, 期望结果：%f", inputHeight, inputWeight, expectedOutput)
	actualOutput, err := CalcBMI(inputHeight, inputWeight)
	if err != nil {
		t.Errorf("expecting no err, but got: %v", err)
	}
	if expectedOutput != actualOutput {
		t.Errorf("expecting %f, but got: %v", expectedOutput, actualOutput)

	}
}
