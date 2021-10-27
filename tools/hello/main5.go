package main

import (
	"fmt"
	"sync"
)

func main() {

	done := make(chan int, 5)

	wg := sync.WaitGroup{}
	// TODO:想想为啥每次done、val的值不同
	// 开N个后台打印线程
	for i := 0; i < cap(done); i++ {
		wg.Add(1)
		go func(){
			done <- i
			fmt.Println("i val:",i)
			wg.Done()
		}()
	}

	// 等待N个后台线程完成
	wg.Wait()


}
