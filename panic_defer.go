package examples_go

import (
	"fmt"
	"io"
	"os"
)

//defer在return之前执行，return不是原子操作;defer后进先出
//返回值 = xxx;调用defer函数;空的return
func f() (result int) {
	defer func() {
		result++
	}()
	return 0
}

func f1() (r int) {
	t := 5
	r = t
	defer func() {
		t = t + 5
		fmt.Println("defer:", t)
	}()
	return
}

func f3() (r int) {
	defer func(r int) {
		r = r + 5 //这里改的r是传值传进去的r，不会改变要返回的那个r值
	}(r)
	return 1
}

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close() //退出执行defer

	return io.Copy(dst, src)
}

func deferOut() {
	for i := 0; i < 4; i++ {
		defer fmt.Print(i) //3210
	}
}

//panic,defer,recover
func fRecover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}
