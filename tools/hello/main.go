package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// 互斥锁线程同步
	var lock sync.Mutex
	lock.Lock()
	go func(){
		fmt.Println("hello goroutine!")
		time.Sleep(1)
		lock.Unlock()
	}()

	// 会阻塞等待unlock之后才执行
	lock.Lock()
	fmt.Println("done")
}
