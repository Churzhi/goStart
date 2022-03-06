package main

import (
	"fmt"
	"os"
)

func main() {
	filePath := "C:\\Users\\Churzhi\\go\\src\\learn.go\\chapter07\\chapter08\\zhou.new"
	writeFile(filePath)

}

func writeFile(filePath string) {
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("无法打开文件", filePath, "错误信息是：", err)
		os.Exit(1) //程序运行失败，直接退出，异常退出
	}
	defer file.Close()

	// handle error
	defer file.Close()
	_, err = file.Write([]byte("this is first line\n"))
	fmt.Println(err)
	// handle error
	_, err = file.WriteString("第二行内容\n")
	fmt.Println(err)
	// handle error
	_, err = file.WriteAt([]byte("CHANGED"), 0)
	// handle error
	fmt.Println(err)
}

// todo 完善append
func writeFileAppend(filePath string) {
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("无法打开文件", filePath, "错误信息是：", err)
		os.Exit(1) //程序运行失败，直接退出，异常退出
	}
	defer file.Close()

	// handle error
	defer file.Close()
	_, err = file.Write([]byte("this is first line\n"))
	fmt.Println(err)
	// handle error
	_, err = file.WriteString("第二行内容\n")
	fmt.Println(err)
	// handle error
	_, err = file.WriteAt([]byte("CHANGED"), 0)
	// handle error
	fmt.Println(err)
}
