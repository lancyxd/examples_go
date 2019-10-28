package examples_go

import (
	"examples_go/demo_kafka/common/logging"
	"fmt"
	"github.com/couchbase/gocb"
)

var (
	Client   *gocb.Cluster //客户端
	bucketDB *gocb.Bucket  //具体的bucket
)

func InitCouchDB(url string) (*gocb.Cluster, error) {
	Client, err := gocb.Connect(url)
	if err != nil {
		return nil, err
	}
	return Client, nil
}

//insert
func addCouchDB(id uint64, info map[string]interface{}) int {
	query := gocb.NewN1qlQuery("INSERT INTO `test-sample` (KEY, VALUE) VALUES ($1,$2)")
	query.AdHoc(false)

	key := fmt.Sprintf("%s%d", "prefix_", id)
	para := []interface{}{key, info}

	_, err := bucketDB.ExecuteN1qlQuery(query, para)
	if err != nil {
		logging.Error("AddCouchDB ExecuteN1qlQuery failed, err=%s, para[key=%s, value=%+v]", err.Error(), key, info)
		return -1
	}
	return 0
}

//get
func getCouchDB(id string) int {
	query := gocb.NewN1qlQuery("SELECT name FROM `test-sample` WHERE id=$1")
	query.AdHoc(false)
	para := []interface{}{id}
	_, err := bucketDB.ExecuteN1qlQuery(query, para)
	if err != nil {
		logging.Error("getCouchDB ExecuteN1qlQuery failed,err=%s\n", err.Error())
		return -1
	}
	logging.Info("getCouchDB ExecuteN1qlQuery sucess!")
	return 0
}

//del
func delCouchDB(id string) int {
	query := gocb.NewN1qlQuery("DELETE FROM `test-sample`  WHERE id = $1")
	query.AdHoc(false)
	para := []interface{}{id}
	_, err := bucketDB.ExecuteN1qlQuery(query, para)
	if err != nil {
		logging.Error("delCouchDB ExecuteN1qlQuery failed,err=%s\n", err.Error())
		return -1
	}
	logging.Info("delCouchDB ok,id=%s", id)
	return 0
}

//update
func updateCouchDB(name, id string) int {
	query := gocb.NewN1qlQuery("UPDATE `test-sample` p SET p.name=$1 Where p.id= $2")
	query.AdHoc(false)
	para := []interface{}{name, id}
	_, err := bucketDB.ExecuteN1qlQuery(query, para)
	if err != nil {
		logging.Error("updateCouchDB ExecuteN1qlQuery failed,err=%s\n", err.Error())
		return -1
	}
	logging.Info("updateCouchDB ok,name=%s,id=%s", name, id)
	return 0
}
