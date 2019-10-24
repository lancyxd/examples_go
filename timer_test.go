package examples_go

import (
	"fmt"
	"testing"
	"time"
)

func writereload() {
	fmt.Println("writereload 执行,time:", time.Now().Format("2006-01-02 15:04:05"))
}

func TestTimerMinute(t *testing.T) {
	minuteTimer(writereload)
	time.Sleep(4 * time.Minute)
}

func TestTimerZero(t *testing.T) {
	zeroTimer(writereload)
	time.Sleep(25 * time.Hour)
}

func TestTimerDouble(t *testing.T) {
	doubleTimer()
	time.Sleep(2 * time.Second)
}
