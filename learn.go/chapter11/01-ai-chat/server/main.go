package main

import (
	"fmt"
	"log"
	"net"
)

var qa = map[string]string{
	"你好": "你好",
}

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		// handle error
		log.Fatal(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
			log.Printf("warning! 建立连接失败:", err)
			continue
		}
		fmt.Println(conn)
	}
}

func talk(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		valid, err := conn.Read(buf)
		if err != nil {
			log.Println("WARNING: 读取数据失败:", err)
			continue
		}
		content := buf[valid]
		resp, ok := qa[string(content)]
		if !ok {
			log.Println("没有找到回答")
		}
	}
}
