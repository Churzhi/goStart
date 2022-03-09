package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"learn.go/pkg/apis"
	"log"
	"os"
)

func NewRecord(filePath string) *record {
	return &record{
		filePath: filePath,
	}
}

type record struct {
	filePath string
}

func (r *record) writeFileWithAppendJson(data []byte) error {
	file, err := os.OpenFile(r.filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println("无法打开文件：", r.filePath, "错误信息是：", err)
		os.Exit(1)
	}
	defer file.Close()
	_, err = file.Write(append(data, '\n'))
	if err != nil {
		return err
	}
	return nil
}

func (r *record) saveInformation(information *apis.PersonalInformation) error {
	data, err := json.Marshal(information)
	if err != nil {
		fmt.Println("marshal 出错：", err)
		return err
	}
	if err := r.writeFileWithAppendJson(data); err != nil {
		log.Println("写入 JSON 时出错：", err)
		return err
	}
	return nil

}

func writeFile(filePath string, data []byte) {
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("无法打开文件", filePath, "错误信息是：", err)
		os.Exit(1)
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func readFile(filePath string) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("读取文件失败：", err)
		return
	}
	fmt.Println("读取出来的内容是：", string(data))
	personalInformation := apis.PersonalInformation{}
	err = json.Unmarshal(data, &personalInformation)
	if err != nil {
		log.Fatal("数据解码失败：", err)
		return
	}

}
