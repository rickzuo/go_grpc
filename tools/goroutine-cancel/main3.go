package main

import (
	"fmt"
	"strconv"
	"time"
)

func Worker3(id int, cancel chan bool, limit chan bool) {
	//defer wg.Done()

	<-limit
	counter ++
	fmt.Println("goid is:", id,"count:",counter)
	for {
		select {
		case <-cancel:
			fmt.Println("quit goid:", id)
			return
		default:
			//fmt.Println("working goid:", id)
		}
	}
}
// 全局变量
var counter int

func main() {

	cancel := make(chan bool)
	// 控制并发数为10
	limit := make(chan bool, 100)
	//var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		//wg.Add(1)
		limit <- true
		fmt.Println("limit len",len(limit))
		go Worker3(i, cancel, limit)
		fmt.Println("===========" + strconv.Itoa(i) + "=============")
	}
	//我们通过close来关闭cancel管道向多个Goroutine广播退出的指令
	close(cancel)
	fmt.Println("done")

	//wg.Wait()
	select {
		case <- time.After(time.Second * 3):
			fmt.Println("shut down:",counter)
	}

}
