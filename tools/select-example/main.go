package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		for {
			select {
			case ch <- 10:
				fmt.Println(10)
			case ch <- 12:
				fmt.Println(12)
			case v := <-ch:
				fmt.Println("v:", v)
			case <-time.After(time.Second * 2):
				fmt.Println("timeout")
				return
			default:
				fmt.Println("default:", <-ch)
			}
		}
	}()

	for v := range ch {
		fmt.Println("v:", v)
	}

}
