package main

import "fmt"

func main() {
	// 有缓冲管道线程同步
	done := make(chan int,1)

	go func() {
		fmt.Println("hello chan")
		//std := <- done
		done <- 10
		//fmt.Println(fmt.Sprintf("std:%d",std))
	}()

	// 使用chan 阻塞 等待返回
	//<-done
	fmt.Println(fmt.Sprintf("done:%#v",<-done))
	//fmt.Println("done",done)
}
