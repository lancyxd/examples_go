//定时器

package examples_go

import (
	"fmt"
	"time"
)

//execute every minute,每隔一分钟执行一次
func minuteTimer(f func()) {
	go func() {
		for {
			t := time.NewTimer(time.Second * 60)
			<-t.C
			f()
		}
	}()

}

//每天零点定时执行函数f
func zeroTimer(f func()) {
	go func() {

		for {
			now := time.Now()                                                                    //2018-11-19 16:33:38.735342 +0800 CST m=+0.007000401
			next := now.Add(time.Hour * 24)                                                      //计算下一个零点
			next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location()) //2018-11-20 00:00:00 +0800 CST

			t := time.NewTimer(next.Sub(now))
			<-t.C
			f() //定时执行的操作
		}

	}()
}

//两个定时器，一个5s 定时执行一次，一个10s定时执行一次
func doubleTimer() {
	t1 := time.NewTimer(time.Second * 5)
	t2 := time.NewTimer(time.Second * 10)

	for {
		select {
		case <-t1.C:
			fmt.Println("5s timer")
			t1.Reset(time.Second * 5)

		case <-t2.C:
			fmt.Println("10s timer")
			t2.Reset(time.Second * 10)
		}
	}
}
