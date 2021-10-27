package main

import "fmt"

func main() {
	// 无缓冲管道线程同步
	done := make(chan int)

	go func() {
		fmt.Println("hello chan")
		<- done
	}()

	// 使用chan 阻塞 等待返回
	done <- -1
	fmt.Println(fmt.Sprintf("done:%#v",done))
}
