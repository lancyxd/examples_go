package main

import (
	"examples_go/demo_thrift/example"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"log"
	"net"
)

const (
	HOST = "127.0.0.1"
	PORT = "8898"
)

func main() {
	tSocket, err := thrift.NewTSocket(net.JoinHostPort(HOST, PORT))
	if err != nil {
		log.Fatal("tSocket error:", err)
	}
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	transport := transportFactory.GetTransport(tSocket)
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	client := example.NewFormatDataClientFactory(transport, protocolFactory)
	if err := transport.Open(); err != nil {
		log.Fatalln("Error opening:", HOST+":"+PORT)
	}
	defer transport.Close()

	data := example.Data{Text: "hello,world!"}
	fmt.Println("data:", data)
	d, err := client.DoFormat(&data)
	fmt.Println(d.Text)
}
