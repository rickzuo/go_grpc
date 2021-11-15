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
	fmt.Println("goid is:", id, "Counter:", Counter)

	// 控制每秒并发数
	time.Sleep(time.Second)
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

	var (
		wg          = sync.WaitGroup{}
		cancel      = make(chan bool)
		timerCancel = make(chan bool)
		// 控制并发数为10
		limit = make(chan struct{}, 3)
	)

	// 每隔1s打印一次结构

	go func() {
		// 注意是ticker 不是 timer := time.NewTimer(time.Second)，这个的话要放到for循环中
		timer := time.NewTicker(time.Second)
		for {
			select {
			case <-timerCancel:
				fmt.Println("timer close:", Counter)
				return
			case <-timer.C:
				fmt.Println("timer clock:", Counter)
			}
			fmt.Println("current token cnt:",time.Now())
		}
	}()


	for i := 0; i < 10; i++ {
		wg.Add(1)
		limit <- struct{}{}
		go Worker2(i, cancel, &wg, limit)
		//fmt.Println("===========" + strconv.Itoa(i) + "=============")
	}
	// 这里block,需要等上面的for循环结束
	fmt.Println("done:", Counter)

	//我们通过close来关闭cancel管道向多个Goroutine广播退出的指令
	close(cancel)
	close(timerCancel)
	wg.Wait()

}
