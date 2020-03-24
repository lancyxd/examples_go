package examples_go

import "fmt"

//func main()  {
////	runtime.GOMAXPROCS(2)
////	var i int
////	go func() {
////		for {
////			i=2
////			//fmt.Println("ffffffffffffffffff")
////			//time.Sleep(time.Millisecond * 20)
////		}
////	}()
////
////
////	for{
////		fmt.Println("hello ,i:",i)
////		time.Sleep(1*time.Second)
////	}
////}



/*
导致死锁的原因:
var chanArr = make([]chan int, 2)仅初始化了chan数组，数组中每个chan没初始化，值为nil；此时无论往通道里面写数据还是读数据均无法进行。
*/


//var chanArr=make([]chan  int,2)
var chanArr =[]chan int{
	make(chan int,2), //初始化，可存放两个int类型
	make(chan int),
}

func run()  {
	chanArr[0]<-1
	chanArr[0]<-2
}

func main()  {
	//chanArrRead()
	chanArrRead1()
}

func chanArrRead()  {
	go run()
	obj1:=<-chanArr[0]
	obj2:=<-chanArr[0]
	fmt.Println(obj1,obj2)
}

func chanArrRead1()  {
	go run()
	obj1:=<-chanArr[0]
	obj2:=<-chanArr[0]
	fmt.Println(obj1,obj2)

	var chanArr [2]chan int
	for i:=range chanArr{
		chanArr[i]=make(chan int,2)
	}
}