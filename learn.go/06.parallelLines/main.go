package main

import (
	"fmt"
)

/*
判断两条线是否平行
提示：
两点决定一条直线
两条线是否平行取决于两条线的斜率是否一样
*/

func main() {

	type Line struct {
		x1        float64
		y1        float64
		x2        float64
		y2        float64
		gradient  float64
		intercept float64
	}

	var line Line
	lines := [2]Line{line}

	for i := 0; i < 2; i++ {
		fmt.Println("请输入直线", i+1, "的两个（X，Y）坐标：")
		line.x1, line.y1, line.x2, line.y2 = getLineCoordinate()
		line.gradient, line.intercept = calcGradientAndIntercept(line.x1, line.y1, line.x2, line.y2)
		lines[i] = line
	}
	//fmt.Println("line:", lineCoordinates)

	for i := 0; i < len(lines); i++ {
		for j := i + 1; j < len(lines); j++ {
			if lines[i].gradient == lines[j].gradient && lines[i].intercept != lines[j].intercept {
				fmt.Println("直线", i+1, "和直线", i+2, "平行")
			} else if lines[i].gradient == lines[j].gradient && lines[i].intercept == lines[j].intercept {
				fmt.Println("直线", i+1, "和直线", i+2, "重合")
			} else if lines[i].gradient*lines[j].gradient == -1 {
				fmt.Println("直线", i+1, "和直线", i+2, "垂直")
			} else {
				fmt.Println("直线", i+1, "和直线", i+2, "不平行")

			}
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

	gradient = (y2 - y1) / (x2 - x1)
	intercept = y1 - (gradient * x1)
	return gradient, intercept
}
