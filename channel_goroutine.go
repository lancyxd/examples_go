package examples_go

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var (
	chanInt         chan int
	nobufferChanInt chan int

	chanStr     chan string
	chanFloat64 chan float64
)

//类似IO多路复用，可以同时监听多个channel的消息状态
func chanSelect() {
	chanStr = make(chan string, 1)
	chanFloat64 = make(chan float64, 1)
	chanStr <- "hello"
	chanFloat64 <- 1.0

	select {
	case value := <-chanStr:
		fmt.Println("chanStr value:", value)
	case value := <-chanFloat64:
		fmt.Println("chanFloat64 value:", value)
		//panic(value)
	default:
		fmt.Println("default excute")
	}
}

//range遍历
func chanRange() {
	//ch := make(chan int, 10)
	ch := make(chan int) //创建一个无缓存channel

	//新建一个goroutine
	go func() {
		for i := 0; i < 2; i++ {
			ch <- i //往通道写数据
		}

		close(ch) //不需要再写数据时，关闭channel
		//ch <- 666 //关闭channel后无法再发送数据

	}()

	/*
		for num := range ch {
			fmt.Println("num = ", num)
		}
	*/

	for {
		num, ok := <-ch
		if !ok {
			break
		}

		fmt.Println("num = ", num)
	}
}

//time超时
func chanTimeOut() {
	chStr := make(chan string, 1)

	//go doTask(ch) finish task while send msg to ch
	go func() {
		time.Sleep(time.Second * 1)
		chStr <- "result 1"
	}()

	timeout := time.After(time.Second * 2)
	select {
	case value := <-chStr:
		fmt.Println("chanTimeOut value:", value)
	case <-timeout:
		fmt.Println("task timeout")
	}
}

/*
//一些 worker goroutine 需要一直循环处理信息，直到收到 quit 信号
msgCh := make(chan struct{})
quitCh := make(chan struct{})
for {
select {
case <- msgCh:
doWork()
case <- quitCh:
finish()
return
}
*/

//sync包中的WaitGroup能够一直等到所有的goroutine执行完成，并且阻塞主线程的执行，直到所有的goroutine执行完成。
//其三个方法为Add(delta int) 添加或减少等待goroutine的数量;Done()等价于Add(-1);Wait() 执行阻塞，直到所有的WaitGroup数量变成0。
func syncWaitGroup() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func(n int) {
			defer wg.Done()
			EchoNumber(n)
		}(i)
	}

	wg.Wait()
}

func EchoNumber(i int) {
	time.Sleep(3e9)
	fmt.Println("EchoNumber i:", i)
}

//GOMAXPROCS and sync WaitGroup
func gomaxprocxSyncWaitGroup() {
	bTime := time.Now()
	runtime.GOMAXPROCS(1) //GOMAXPROCS设置为1影响goroutine并发，后续代码的go func()相当于串行

	var wg sync.WaitGroup
	wg.Add(20)

	//go协程
	for i := 0; i < 10; i++ { //i为外部for循环的变量，地址不变化。遍历完成后，最终i为10。
		go func() {
			fmt.Println("for_out i: ", i)
			wg.Done()
		}()
	}

	//go协程
	for i := 0; i < 10; i++ { //go func中i为函数参数，与外部for中的i是两个变量，尾部i发生值拷贝，go func内部指向值拷贝的地址
		go func(i int) {
			fmt.Println("i: ", i)
			wg.Done()
		}(i)
	}

	wg.Wait()
	eTime := time.Now()
	fmt.Println("gomaxprocxSyncWaitGroup time cost :", eTime.Sub(bTime).Seconds())
}

var wgro sync.WaitGroup

func syncGoroutine() {
	bTime := time.Now()
	runtime.GOMAXPROCS(runtime.NumCPU()) //设置最大的可同时使用线程数

	for i := 0; i < 10; i++ {
		wgro.Add(1) //每创建一个goroutine，就把任务队列中任务的数量+1
		go calcute(i)
	}

	wgro.Wait() //Wait()这里会发生阻塞，直到队列中所有的任务结束就会解除阻塞
	eTime := time.Now()
	fmt.Println("syncGoroutine time cost :", eTime.Sub(bTime).Seconds())

}

/*多线程比单线程快，时间加速明显*/
func calcute(num int) {
	for i := 1; i < 1000000000; i++ {
		num = num + i
		num = num - i
		num = num * i
		num = num / i
	}
	wgro.Done()
}
