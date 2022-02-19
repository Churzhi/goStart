package main

import (
	"fmt"
	"sync"
	"time"
)

type WebServerV1 struct {
	config     Config
	configLock sync.RWMutex
}

func (ws *WebServerV1) reload() {
	ws.configLock.RLock()
	defer ws.configLock.RUnlock()
	ws.config.Content = fmt.Sprintf("%s", time.Now())
}

func (ws *WebServerV1) ReloadWorker() {
	for {
		time.Sleep(10 * time.Microsecond)
		ws.reload()
	}
}

func (ws *WebServerV1) Visit() string {
	ws.configLock.RLock()
	defer ws.configLock.RUnlock()
	return ws.config.Content
}
