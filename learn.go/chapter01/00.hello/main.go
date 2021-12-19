package main

import "fmt"

func main() {
	fmt.Println("1. 打开冰箱")
	fmt.Println("2. 把大象装进去")
	fmt.Println("3. 关上冰箱")
	var (
		touHigh  = 6
		jianHigh = 25
		kuang    = 50
	)
	for i := 1; i <= touHigh+jianHigh; i++ {
		for j := 1; j <= kuang; j++ {
			//上三角
			if i <= touHigh {
				if j >= (kuang/2+1)+1-i && j <= (kuang/2+1)-1+i {
					fmt.Print("*")
				} else {
					fmt.Print("-")
				}

			}
			//上三角一下部分
			if i > touHigh && i <= jianHigh {
				if j >= (kuang/2+1)+1-i && j <= kuang-3*(i-touHigh) {
					fmt.Print("*")
				} else if j <= (kuang/2+1)-1+i && j >= 0+3*(i-touHigh) {
					fmt.Print("*")
				} else {
					fmt.Print("-")
				}
			}
			if j == kuang {
				fmt.Print("\n")
			}
		}

	}

}
