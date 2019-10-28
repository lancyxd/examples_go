package examples_go

import (
	"fmt"
	"testing"
)

func TestChanInt(t *testing.T) {
	chanInt = make(chan int, 1)      //有缓冲chan int
	nobufferChanInt = make(chan int) //无缓冲chan int

	chanInt <- 1
	fmt.Println("TestChanInt read chan data:", <-chanInt)

	nobufferChanInt <- 1 //无缓冲，写入数据会导致死锁
	fmt.Println("TestChanInt read no buffer chan data:", <-nobufferChanInt)
}

func TestChanWrite(t *testing.T) {
	chanSelect()
}

func TestChanRange(t *testing.T) {
	chanRange() //num =  0;num =  1
}

func TestChanTimeout(t *testing.T) {
	chanTimeOut()
}

func TestSyncWaitGroup(t *testing.T) {
	syncWaitGroup()
}

func TestGomaxprocx(t *testing.T) {
	gomaxprocxSyncWaitGroup()
}

func TestSyncGoroutine(t *testing.T) {
	syncGoroutine()
}
