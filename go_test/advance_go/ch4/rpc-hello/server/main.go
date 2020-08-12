package main

import (
	"log"
	"net"
	"net/rpc"
)

type HelloService struct{}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

func (p *HelloService) Bye(request string, reply *string) error {
	*reply = "Bye:" + request
	return nil
}

func main() {
	// register name and handler , 所有的注册的方法都会在HelloService这个服务空间下
	// 该方式可以注册HelloService这个type下的所有方法
	rpc.RegisterName("HelloService", new(HelloService))
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error：", err)
	}

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error:", err)
	}

	rpc.ServeConn(conn)
}
