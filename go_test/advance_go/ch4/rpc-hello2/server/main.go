package main

import (
	"log"
	"net"
	"net/rpc"
)

const HelloServiceName = "path/to/pkg.HelloService"

// 通过interface实现多态，能够兼容多种rpc service
type HelloServiceInterface interface {
	Hello(request string, reply *string) error
}

func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, svc)
}

type Hello struct {
}

// implement the interface
func (p Hello) Hello(request string, reply *string) error {
	*reply = "Hello:" + request
	return nil
}

func main() {
	var test Hello
	RegisterHelloService(test)

	lister, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := lister.Accept()
	if err != nil {
		log.Fatal(err)
	}

	rpc.ServeConn(conn)
}
