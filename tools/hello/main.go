package main

import (
	"fmt"
	"sync"
)

func main() {
	var lock sync.Mutex

	go func(){
		fmt.Println("hello goroutine!")
		lock.Lock()
	}()

	lock.Unlock()

}
