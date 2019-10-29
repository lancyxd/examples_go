package demo_es

import (
	"examples_go/demo_kafka/common/logging"
	"testing"
)

func TestInitES(t *testing.T) {
	name := "xxxx"
	passwd := "xxxxxx"
	urlList := []string{"http://xxx.xxx.xxx.xxx:xxxx"}

	err := InitES(urlList, name, passwd)
	if err != nil {
		logging.Error("TestInitES InitES failed,err=%s", err.Error())
		return
	}
	logging.Info("TestInitES InitES ok,urls=%+v", urlList)
}

func TestQuery(t *testing.T) {
	//init es
	name := "xxxx"
	passwd := "xxxxxx"
	urlList := []string{"http://xxx.xxx.xxx.xxx:xxxx"}

	err := InitES(urlList, name, passwd)
	if err != nil {
		logging.Error("TestInitES InitES failed,err=%s", err.Error())
		return
	}
	logging.Info("TestInitES InitES ok,urls=%+v", urlList)

	//query
	queryList := QueryList{
		ItemArray: []queryItem{{field: "createTime"}, {field: "content", text: "问题", boost: 3.0, value: "20"}},
		fieldAggs: "colors",
		fieldTime: "create_time",
		Scale:     "1d",
		Offset:    "30d",
		BoostMode: "sum",
		timeFrom:  "2018-10-28",
		timeTo:    "2018-11-28",
		From:      0,
		Size:      10,
		decay:     0.5,
	}

	maplist, err := queryList.Query()
	if err != nil {
		logging.Error("TestQuery queryList.Query fail,err=%s", err.Error())
		return
	}
	logging.Info("TestQuery queryList.Query ok,maplist=%+v", maplist)
}
