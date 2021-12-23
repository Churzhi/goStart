package main

import "testing"

// 测试案例

func TestCase1(t *testing.T) {
	{
		inputRecord("王强", 0.38)
		inputRecord("王强", 0.32)
		randOfWQ, fatRateOfWQ := getRank("王强")
		if randOfWQ != 1 {
			t.Fatalf("预期：王强 第一、但是得到的是：%d", randOfWQ)

		}
		if fatRateOfWQ != 0.32 {
			t.Fatalf("预期：王强 第一、但是得到的是：%f", fatRateOfWQ)

		}
	}

	{
		inputRecord("李静", 0.38)
		inputRecord("李静", 0.32)
		randOfWLJ, fatRateOfLJ := getRank("王强")
		if randOfWLJ != 1 {
			t.Fatalf("预期：李静 第一、但是得到的是：%d", randOfWLJ)

		}
		if fatRateOfLJ != 0.32 {
			t.Fatalf("预期：李静 第一、但是得到的是：%f", fatRateOfLJ)

		}
	}

}
