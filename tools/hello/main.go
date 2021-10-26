package main

import (
	"fmt"
	"sync"
)

func main() {
	var lock sync.Mutex

	go func(){
		lock.Lock()
		fmt.Println("hello goroutine!")

		lock.Unlock()
	}()

	select {

	}


}
