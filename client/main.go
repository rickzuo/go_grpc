package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:9090")
	if err != nil {
		log.Fatalf("client conn fail:%s", err)
	}

	var rsp string
	err = client.Call("HelloService.SayHello", "hello-x", &rsp)
	if err != nil {
		log.Fatalf("call fail:%s", err)
	}

	fmt.Printf("reply:%s \n", rsp)

}
