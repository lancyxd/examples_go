package examples_go

import (
	"examples_go/demo_kafka/common/logging"
	"testing"
)

func TestCouchDB(t *testing.T) {
	url := "http://myserver:8091/" //"http://test-sample:123456@myserver:8091/"
	var err error
	Client, err = InitCouchDB(url)
	if err != nil {
		logging.Error("TestCouchDB  InitCouchDB err,err=%s", err.Error())
		return
	}
	logging.Info("TestCouchDB InitCouchDB ok,url=%s", url)

	bucketDB, err = Client.OpenBucket("test-sample", "pwd123")
	if err != nil {
		logging.Error("TestCouchDB OpenBucket  err,err=%s", err)
		return
	}

	//insert
	info := map[string]interface{}{
		"id":       1,
		"name":     "Lily",
		"age":      5,
		"subjects": `["Math","English","Chinese"]`,
		"geo": `{
			"accuracy": "ROOFTOP",
			"lat": 37.7825,
			"lon": -122.393
		}`,
	}
	errCode := addCouchDB(1, info)
	if errCode < 0 {
		logging.Error("AddCouchDB failed")
		return
	}
	logging.Info("AddCouchDB ok!")
}
