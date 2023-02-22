package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")
	go process1()
	time.Sleep(100 * time.Millisecond)
	// 匿名函数
	go func() {
		fmt.Println("test")
	}()
	fmt.Println("end")
}

func process1() {
	fmt.Println("processing")
}
