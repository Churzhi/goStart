package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	filePath := "C:\\Users\\Churzhi\\go\\src\\learn.go\\chapter07\\chapter08\\zhou"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("无法打开文件", filePath, "错误信息是：", err)
		os.Exit(1) //程序运行失败，直接退出，异常退出
	}
	defer file.Close()

	b := make([]byte, 1024)
	for i := 0; i < 2; i++ {
		n, err := file.Read(b)
		if err != nil {
			fmt.Println("无法读取文件", err)
			os.Exit(2)
		}
		fmt.Println("读出的内容：", string(b)) //如果不转换，则输出的是 byte 类型内容
		fmt.Println("n的大小：", n)
		fmt.Println("读出的内容：", string(b[:n])) //后续程序使用该数据时，需要截取到实际读取的数据，而不是全部
		file.Seek(0, io.SeekStart)           //SeekStart表示从文件头开始数，转移游标
		file.Seek(3, io.SeekCurrent)         //SeekCurrent表示从目前游标位置开始数，转移相对位置

	}

}
