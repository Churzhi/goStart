package main

import "fmt"

func main() {

	//低效
	names := []string{"xiaozhou", "xiaozhang"}
	fr := []float64{28, 18, 18}
	names = append(names, "xiaoli")
	fr = append(fr, 19)

	for i, name := range names {
		if name == "xiaozhou" {
			fmt.Printf("%s的体脂率是%f\n", name, fr[i])
		}
	}
	//如果某个元素具有唯一性
	// Map 定义
	var m1 map[string]int = nil
	//m1["a"] = 1 // panic on nil map
	delete(m1, "a")
	fmt.Println("m1 没有实例化，但是能读", m1["a"])

	m2 := map[string]int{}
	m3 := map[string]int{"王强": 60, "李静": 83, "苗文": 91}
	fmt.Println(m1, m2, m3)

	// 查看是否拿到假的Map键值
	// 判断ok返回值是否是真
	xzScore1, ok1 := m2["王强"]
	fmt.Println(xzScore1, ">>>>", ok1)
	m2["王强"] = 88
	xzScore2, ok2 := m2["王强"]
	fmt.Println(xzScore2, ">>>>", ok2)
	//遍历
	for k, v := range m3 {
		fmt.Printf("%s\t%d\n", k, v)
	}

	var a1 map[int][]string = map[int][]string{11: []string{"苹果", "香蕉"}}
	fmt.Println(a1)
	a2 := map[string]map[int]float64{}
	fmt.Println(a2)

	//增删改查
	m3["小强"] = 77
	fmt.Println(m3)
	delete(m3, "王强")
	fmt.Println(m3)
	m3["小强"] = 80
	fmt.Println(m3)
	fmt.Println(m3["小强"])

	// 难读 无限嵌套
	mSurprise := map[string]map[string]map[float64]int{}
	fmt.Println(mSurprise)

}
