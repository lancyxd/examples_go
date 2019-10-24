package main

import (
	"examples_go/demo_proto/proto"
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"os"
)

func PhoneType(v test.PhoneType) *test.PhoneType {
	return &v
}

func write() {
	p1 := &test.Person{
		Id:   proto.Int32(1),
		Name: proto.String("小张"),
		Phones: []*test.Phone{
			{Type: PhoneType(test.PhoneType_HOME), Number: proto.String("111111111")},
			{Type: PhoneType(test.PhoneType_WORK), Number: proto.String("222222222")},
		},
	}

	p2 := &test.Person{
		Id:   proto.Int32(2),
		Name: proto.String("小王"),
		Phones: []*test.Phone{
			{Type: PhoneType(test.PhoneType_HOME), Number: proto.String("333333333")},
			{Type: PhoneType(test.PhoneType_WORK), Number: proto.String("444444444")},
		},
	}

	//创建地址薄
	book := &test.ContactBook{}
	book.Persons = append(book.Persons, p1)
	book.Persons = append(book.Persons, p2)

	//编码数据
	data, _ := proto.Marshal(book)
	fmt.Printf("%+v\n", string(data))
	ioutil.WriteFile("./test.txt", data, os.ModePerm)
}

func read() {
	data, _ := ioutil.ReadFile("./test.txt")
	book := &test.ContactBook{}
	proto.Unmarshal(data, book)

	for _, v := range book.Persons {
		fmt.Println(v.Id, v.Name)
		for _, vv := range v.Phones {
			fmt.Println(vv.Type, vv.Number)
		}
	}

}

func main() {
	write()
	read()
}
