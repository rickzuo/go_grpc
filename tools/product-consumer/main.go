package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func producer(factor int, ch chan<- int) {

	for i := 0; ; i++ {
		data := i * factor
		ch <- data
		fmt.Println("product:", data)
	}
}

func consumer(ch <-chan int) {
	for v := range ch {
		fmt.Println(fmt.Sprintf("consumer:%d", v))
	}
}

func main() {
	fmt.Println("start work")
	// 如果ch满了会报错
	ch := make(chan int, 10)
	go producer(3, ch)
	go producer(5, ch)
	go consumer(ch)


	fmt.Println("done")
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println(fmt.Sprintf("quit :(%v)", <-sig))
}
