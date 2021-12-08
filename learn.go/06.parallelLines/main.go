package main

import (
	"fmt"
)

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
	//fmt.Println("coordinate:", lineCoordinates)

	for i := 0; i < len(lineCoordinates); i++ {
		for j := i + 1; j < len(lineCoordinates); j++ {
			if lineCoordinates[i].gradient == lineCoordinates[j].gradient && lineCoordinates[i].intercept != lineCoordinates[j].intercept {
				fmt.Println("直线", i+1, "和直线", i+2, "平行")
			} else if lineCoordinates[i].gradient == lineCoordinates[j].gradient && lineCoordinates[i].intercept == lineCoordinates[j].intercept {
				fmt.Println("直线", i+1, "和直线", i+2, "重合")
			} else if lineCoordinates[i].gradient*lineCoordinates[j].gradient == -1 {
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
