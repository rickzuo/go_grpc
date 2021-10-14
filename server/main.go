package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

type HelloService struct {
}

// rpc规定：方法只能有两个可序列化的参数，其中第二个必须要是指针类型，并且返回一个err,不能返回其他。如果第二个参数不是指针类型会抛出如下错误，可以自己尝试下！
//2021/10/14 08:58:45 rpc.Register: reply type of method "SayHello" is not a pointer: "string"
//2021/10/14 08:58:45 rpc.Register: type HelloService has no exported methods of suitable type
//2021/10/14 08:58:45 regsiter rpc server fail:rpc.Register: type HelloService has no exported methods of suitable type

func (service *HelloService) SayHello(req string, reply *string) error {

	*reply = fmt.Sprintf("hello %s", req)
	fmt.Println(reply)
	return nil
}

func main() {
	// 将helloservice注册为rpc服务
	err := rpc.RegisterName("HelloService", new(HelloService))
	if err != nil {
		log.Fatalf("regsiter rpc server fail:%s\n", err)
	}

	listen, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("Listen err:%s", err)
	}
	conn, err := listen.Accept()
	if err != nil {
		log.Fatalf("accept err:%s", err)
	}

	rpc.ServeConn(conn)

}
