package main

import "testing"

func TestGetHealthinessSuggestions(t *testing.T) {
	expected := [3]string{"目前是：标准", "目前是：标准", "目前是：标准"}
	person := inputRecordFake()
	suggestion, err := getHealthinessSuggestions(person)
	t.Log("fatRate:", suggestion)
	if err == nil {
		// 为什们 err 不为空的时候 测试不通过的？
		t.Errorf("")
	}
	for i, item := range expected {
		if suggestion[i] != item {
			t.Errorf("expected is different from suggestion")
		}
	}

}
