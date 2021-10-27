package main

import "fmt"

func main() {

	done := make(chan int, 5)

	// TODO:想想为啥每次done、val的值不同
	// 开N个后台打印线程
	for i := 0; i < cap(done); i++ {
		go func(){
			done <- i
			fmt.Println("i val:",i)
		}()
	}

	// 等待N个后台线程完成
	for i:=0; i < cap(done);i++{
		fmt.Println("done:",<- done)
	}


}
