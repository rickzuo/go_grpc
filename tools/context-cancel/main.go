package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(ctx context.Context, wg *sync.WaitGroup) error {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("goroutine quit")
			return ctx.Err()
		default:
			fmt.Println("working")
		}
	}
}

func main() {
	// go1.7 after
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)

	defer cancel()
	var wg = sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(ctx, &wg)
	}

	wg.Wait()
}
