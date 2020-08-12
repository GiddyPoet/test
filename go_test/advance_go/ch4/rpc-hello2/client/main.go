package main

import (
	"fmt"
	"log"
	"net/rpc"
)

const HelloSerivceName = "path/to/pkg.HelloService"

type HelloSerivceClient struct {
	*rpc.Client
}

func DialHelloService(network, address string) (*HelloSerivceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloSerivceClient{Client: c}, nil
}

func (p *HelloSerivceClient) Hello(request string, reply *string) error {
	return p.Client.Call(HelloSerivceName+".Hello", request, reply)
}

func main() {
	var reply string
	client, err := DialHelloService("tcp", "localhost:1234")
	if err != nil {
		log.Fatal(err)
	}

	client.Hello("GiddyPoet", &reply)
	fmt.Println(reply)
}
