package main

import (
	"fmt"
	"sync"
	"time"
)

// 互斥锁
var (
	counter = 0
	lock    sync.Mutex
)

func main() {
	for i := 0; i < 20; i++ {
		go incr()
	}

	time.Sleep(10 * time.Microsecond)
}

func incr() {
	lock.Lock()
	defer lock.Unlock()
	counter++
	fmt.Println(counter)
}
