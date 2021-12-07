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

	type Coordinate struct {
		x1        float64
		y1        float64
		x2        float64
		y2        float64
		gradient  float64
		intercept float64
	}

	var coordinate Coordinate
	lineCoordinates := [2]Coordinate{coordinate}

	for i := 0; i < 2; i++ {
		fmt.Println("请输入直线", i+1, "的两个（X，Y）坐标：")
		coordinate.x1, coordinate.y1, coordinate.x2, coordinate.y2 = getLineCoordinate()
		coordinate.gradient, coordinate.intercept = calcGradientAndIntercept(coordinate.x1, coordinate.y1, coordinate.x2, coordinate.y2)
		lineCoordinates[i] = coordinate
	}
	fmt.Println("coordinate:", lineCoordinates)

	for i := 0; i < len(lineCoordinates); i++ {
		if lineCoordinates[i].gradient == lineCoordinates[i+1].gradient && lineCoordinates[i].intercept != lineCoordinates[i+1].intercept {
			fmt.Println("直线", i+1, "和直线", i+2, "平行")
		} else if lineCoordinates[i].gradient == lineCoordinates[i+1].gradient && lineCoordinates[i].intercept == lineCoordinates[i+1].intercept {
			fmt.Println("直线", i+1, "和直线", i+2, "重合")

		} else if lineCoordinates[i].gradient*lineCoordinates[i+1].gradient == -1 {
			fmt.Println("直线", i+1, "和直线", i+2, "垂直")
		} else {
			fmt.Println("直线", i+1, "和直线", i+2, "不平行")

		}
	}
}

func getLineCoordinate() (x1 float64, y1 float64, x2 float64, y2 float64) {

	for {
		fmt.Scanln(&x1, &y1, &x2, &y2)
		if x1 == x2 && y1 == y2 {
			fmt.Println("直线两点不能相同哦，请重新输入：")
		}
		break
	}
	return x1, y1, x2, y2
}

func calcGradientAndIntercept(x1 float64, y1 float64, x2 float64, y2 float64) (gradient float64, intercept float64) {
	// 点斜式：y-y0=k（x-x0），斜率就是k
	// y = kx +b

	/*若直线L1:A1x+B1y+C1 =0与直线 L2:A2x+B2y+C2=0

	1. 当A1B2-A2B1≠0时, 相交

	2.A1/A2=B1/B2≠C1/C2, 平行

	3.A1/A2=B1/B2=C1/C2, 重合

	4.A1A2+B1B2=0, 垂直

	*/
	gradient = (y2 - y1) / (x2 - x1)
	intercept = y1 - (gradient * x1)
	return gradient, intercept
}
