package main

import (
	"fmt"
	"time"
)

func Worker(cancel chan bool) {

	for {
		select {
		case <-cancel:
			fmt.Println("quit")
			return
		default:
			fmt.Println("working")
		}
	}
}
func main() {

	cancel := make(chan bool)
	go Worker(cancel)
	fmt.Println("done")

	for {
		select {
		case <-time.After(time.Second*5):
			fmt.Println("send cancel")
			cancel <- true
			return
		}
	}


}
