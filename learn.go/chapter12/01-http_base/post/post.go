package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	p := strings.NewReader("foooo")
	resp, err := http.Post("https://www.baidu.com", "*/*", p)
	if err != nil {
		log.Fatal("无法发送请求：", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal("Body关闭失败：", err)
		}
	}(resp.Body)
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("无法读取返回内容：", err)
	}
	fmt.Println(data)
}
