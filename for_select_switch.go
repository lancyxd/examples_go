package examples_go

import (
	"fmt"
	"time"
)

var (
	c1 = make(chan string)
	c2 = make(chan string)
)

//select 当channel中c1无数据再读取会导致死锁
func forSelect() {
	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("forSelect received msg1:", msg1)
		case msg2 := <-c2:
			fmt.Println("forSelect received msg2:", msg2)
		}
	}
}

//break跳出的是select,并不是break for
func forSelectBreak(i int) {
	for {
		fmt.Println("forSelectBreak inside for:", i)

		select {
		case c := <-time.After(time.Second * time.Duration(1)):
			fmt.Println(c.Unix())
			i++
			if i == 2 {
				fmt.Println("forSelectBreak break now")
				break
			}
			fmt.Println("forSelectBreak inside the select: ")
		}

	}
}

func forSelectContinue(i int) {
	for {
		fmt.Println("forSelectContinue inside for:", i)

		select {
		case c := <-time.After(time.Second * time.Duration(1)):
			fmt.Println(c.Unix())
			i++
			if i == 2 {
				fmt.Println("forSelectContinue continue now")
				continue
			}
			fmt.Println("forSelectContinue inside the select:")
		}
	}
}

//跳出当前循环jump out of for loop,no longer execute
func forBreak(i int) {
	for {
		fmt.Println("forBreak inside for:", i)
		i++
		if i == 2 {
			fmt.Println("forBreak break now")
			break
		}
	}
}

//跳过本次循环,继续执行接下来的循环
func forContinue(i int) {
	for {
		fmt.Println("forContinue inside for:", i)
		time.Sleep(2 * time.Second)
		i++
		if i == 2 {
			fmt.Println("forContinue continue now")
			continue
		}
	}
}

//switch
func Switch(i int) {
	switch i {
	case 1:
		fmt.Println("Switch one")
	case 2:
		fmt.Println("Switch two")
	case 3:
		fmt.Println("Switch three")
	default:
		fmt.Println("Switch not 1,2,3")
	}
}

//break跳出的是switch,并不是break for
func forSwitchBreak() {
	fmt.Println("forSwitchBreak for outside")
	for {
		fmt.Println("forSwitchBreak inside for:")
		time.Sleep(1 * time.Second)
		switch {
		case true:
			fmt.Println("forSwitchBreak breaking out...")
			break
		}
	}
}

func forSwitchContinue() {
	fmt.Println("forSwitchContinue for outside")
	for {
		fmt.Println("forSwitchContinue inside for:")
		time.Sleep(1 * time.Second)
		switch {
		case true:
			fmt.Println("forSwitchContinue continue out...")
			continue
		}
	}
}

//遇到select和switch,break跳出for循环的方法
func forBreakOut() {
LOOP:
	for {
		switch {
		case true:
			fmt.Println("breaking out...")
			break LOOP
		}
	}

	fmt.Println("out")
}
