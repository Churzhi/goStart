package main

import (
	"fmt"
)

/*
前提：在直角平面坐标系
两点确定一条直线（两点不能重合）
01.录入直线1在X,Y的值
02.录入直线2在X,Y的值
03.录入斜率K1,K2(条件两条直线平行，斜率相等)
04.录入两条直线在Y轴上的点b1,b2(b1不等于了2)
05录入数据都符合，那么这两条直线平行，否则重新录入

*/
func main() {

	type coordinate struct {
		x float64
		y float64
	}
	var (
		lineOneX1 float64
		lineOneY1 float64
		lineOneX2 float64
		lineOneY2 float64
		lineTwoX1 float64
		lineTwoY1 float64
		lineTwoX2 float64
		lineTwoY2 float64

		//斜率
		gradient float64
	)

	fmt.Println("请输入直线1的两个（X，Y）坐标：")
	fmt.Scanln(&lineOneX1, &lineOneY1, &lineOneX2, &lineOneY2)

	fmt.Println("请输入直线2的（X，Y）坐标：")
	fmt.Scanln(&lineTwoX1, &lineTwoY1, &lineTwoX2, &lineTwoY2)

	// 点斜式：y-y0=k（x-x0），斜率就是k
	gradient =
}
