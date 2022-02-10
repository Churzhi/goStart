package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {

}

func account(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

}

func accountRegister(ctx context.Context, doneCh chan string) {
	fmt.Println("注册账号")
	defer fmt.Println("注册账号完成")
	for { // ... 调用xxx接口
		select {
		case <-ctx.Done():
			fmt.Println("context 结束，不再注册")
			return
		default:

		}
		doneCh <- "accountRegister"
		break
	}
}

func accountGrantAuth(ctx context.Context, doneCh chan string) {
	fmt.Println("授权账号")
	defer fmt.Println("授权账号完成")
}

func distributeService(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, 7*time.Minute)
	defer cancel()

	wg := sync.WaitGroup{}
	wg.Add(2)
	go distributeLB(ctx, &wg)
	go distributeInstance(ctx, &wg)
}

\