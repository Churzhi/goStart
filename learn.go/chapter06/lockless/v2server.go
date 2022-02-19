package main

import (
	"fmt"
	"time"
)

type WebServerV2 struct {
	config *Config
	// 单指针读写，原子操作
	// go 的赋值特性
}

func (ws *WebServerV2) reload() {
	// 写回时使用指针
	ws.config = &Config{
		Content: fmt.Sprintf("%d", time.Now().UnixNano()),
	}
}

func (ws *WebServerV2) ReloadWorker() {
	for {
		time.Sleep(10 * time.Microsecond)
		ws.reload()
	}
}

func (ws *WebServerV2) Visit() string {
	return ws.config.Content
}
