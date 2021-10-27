package main

import (
	"fmt"
	"strconv"
	"sync"
)

func Worker2(id int, cancel chan bool, wg *sync.WaitGroup, limit chan bool) {
	defer wg.Done()
	fmt.Println("goid is:", id)
	<-limit

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

func main() {

	cancel := make(chan bool)
	// 控制并发数为10
	limit := make(chan bool, 10)
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		limit <- true
		fmt.Println("limit len",len(limit))
		go Worker2(i, cancel, &wg, limit)
		fmt.Println("===========" + strconv.Itoa(i) + "=============")
	}
	//我们通过close来关闭cancel管道向多个Goroutine广播退出的指令
	close(cancel)
	fmt.Println("done")

	wg.Wait()

}
