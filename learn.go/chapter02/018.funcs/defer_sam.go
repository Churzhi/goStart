package main

import (
	"fmt"
	"os"
	"time"
)

func openFile() {
	fileName := "test.txt"

	fileObj, err := os.Open(fileName)
	if err != nil {
		fmt.Println("error", err)
		os.Exit(1)
	}
	defer fileObj.Close()
}

func deferGuess() {
	startTime := time.Now()
	fmt.Println("start time:", startTime)
	// 运行到这里时，已经准备好所有东西了
	defer fmt.Println("duration:", time.Now().Sub(startTime))

	defer func() {
		//统计时间准确
		fmt.Println("duration upgraded:", time.Now().Sub(startTime))
	}()
	time.Sleep(5 * time.Second)
	fmt.Println("finish time:", time.Now())
}
