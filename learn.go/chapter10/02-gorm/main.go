package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

func connectDB() {
	conn, err := gorm.Open(mysql.Open("root:learngo@tcp(127.0.0.1:3306)/learngo"))
	if err != nil {
		log.Print(err)
	}

	name := fmt.Sprintf("orm-%s", time.Now().Format(time.RFC3339))

}
func createNewPerson(conn *gorm.DB) error {
	resp := conn.Create(&PersonalInformation{
		Name:   "zhou",
		Sex:    "男",
		Tall:   1.8,
		Weight: 90,
		Age:    23,
	})
	if err := resp.Error; err != nil {
		//TODO处理错误信息
		fmt.Println("创建***时失败：", err)
		return err
	}
	fmt.Println("创建***成功：")
	return nil
}

func ormSelect(conn *gorm.DB)  {
	output := make([]*PersonalInformation, 0)
	resp := conn.Where(&PersonalInformation{Name: "zhou"}).Find(&output)
	if resp.Error != nil {
		log.Printf("查找过程中出错：%v", resp.Error)
		// TODO handle error
		fmt.Println("查找失败：",resp.Error)
		return
	}
}

func updateExistingPerson(conn *gorm.DB) error  {
	// 从数据库获取已有的数据
	var person *PersonalInformation = queryFromDatabase("小强")
	// 修改最近的身体变化
	person.Tall = 1.91
	person.Weight = 95
	person.Age = 27
	resp := conn.Updates(person) // 全量保存
	if resp.Error != nil {
		// TODO handle error
	}
}
func main() {
	conn := connectDB
	fmt.Println(conn)
	err : =createNewPerson(conn)
	if err:=
}
