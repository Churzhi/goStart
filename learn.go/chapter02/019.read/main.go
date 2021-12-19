package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"learn.go.tools/calc"
)

func main() {

	//录入
	var (
		name   string
		sex    string
		tall   float64
		weight float64
		age    int
	)
	//arguments := os.Args
	//fmt.Println(arguments)

	cmd := &cobra.Command{
		Use:   "healthCheck",
		Short: "体脂计算器",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("name:", name)
			fmt.Println("sex:", sex)
			fmt.Println("tall:", tall)
			fmt.Println("weight:", weight)
			fmt.Println("age:", age)
			// 计算
			// 使用 learn.go.tools
			bmi := calc.CalcBMI(tall, weight)

			fatRate := calc.CalcFatRate(bmi, age, sex)

			// 评估结果
			fmt.Println("fatRate:", fatRate)
		},
	}

	cmd.Flags().StringVar(&name, "name", "", "姓名")
	cmd.Flags().StringVar(&sex, "sex", "", "性别")
	cmd.Flags().Float64Var(&tall, "tall", 0.0, "身高")
	cmd.Flags().Float64Var(&weight, "weight", 0.0, "身高")
	cmd.Flags().IntVar(&age, "age", 0, "身高")

	// 运行自定义的命令行对象
	cmd.Execute()
}
