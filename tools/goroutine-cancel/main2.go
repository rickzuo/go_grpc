package main

import (
	"fmt"
	"sync"
	"time"
)

var Counter = 0

func Worker2(id int, cancel chan bool, wg *sync.WaitGroup, limit chan struct{}) {
	defer wg.Done()

	Counter++
	time.Sleep(time.Second)
	fmt.Println("goid is:", id, "Counter:", Counter)
	// 注意这个位置
	<-limit
	select {
		case <-cancel:
			fmt.Println("quit goid:", id)
			return
		default:
		//fmt.Println("working goid:", id)
	}
}

func main() {

	cancel := make(chan bool)
	// 控制并发数为10
	limit := make(chan struct{}, 3)
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		limit <- struct{}{}
		go Worker2(i, cancel, &wg, limit)
		//fmt.Println("===========" + strconv.Itoa(i) + "=============")
	}
	//我们通过close来关闭cancel管道向多个Goroutine广播退出的指令
	close(cancel)
	fmt.Println("done:", Counter)

	wg.Wait()

}
