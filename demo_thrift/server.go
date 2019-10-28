package main

import (
	"examples_go/demo_thrift/example"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"log"
	"strings"
)

type FormatDataImpl struct{}

func (fdi *FormatDataImpl) DoFormat(data *example.Data) (r *example.Data, err error) {
	var rData example.Data
	rData.Text = strings.ToUpper(data.Text)

	return &rData, nil
}

const (
	HOST = "127.0.0.1"
	PORT = "8898"
)

func main() {
	handler := &FormatDataImpl{}
	processor := example.NewFormatDataProcessor(handler) //先建立和服务器的连接的socket，再通过socket建立Transport
	serverTransport, err := thrift.NewTServerSocket(HOST + ":" + PORT)
	if err != nil {
		log.Fatalln("Error:", err)
	}
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault() //创建二进制协议

	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	fmt.Println("Running at:", HOST+":"+PORT)
	server.Serve()
}
