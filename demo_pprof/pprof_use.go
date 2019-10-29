package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

/*
imoport导包
添加应用进程
*/

func main() {

	go func() {
		for {
			log.Println(Add("https://github.com/EDDYCJY"))
		}
	}()

	http.ListenAndServe("0.0.0.0:8082", nil)
}

var datas []string

func Add(str string) string {
	data := []byte(str)
	sData := string(data)
	datas = append(datas, sData)
	return sData
}
