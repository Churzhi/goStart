package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
)

func main() {

	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	learnDB, err := sql.Open("mysql", "root:chu8837552@tcp(127.0.0.1:3306)/test_01")
	if err != nil {
		fmt.Println("链接数据库失败:", err)
	}
	defer learnDB.Close()
	if err = learnDB.Ping(); nil != err {
		fmt.Println("DB 测试失败:", err)
	}
	fmt.Println("链接数据库成功:")

}
