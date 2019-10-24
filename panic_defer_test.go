package examples_go

import (
	"fmt"
	"testing"
)

func TestDefer(t *testing.T) {
	fmt.Println(f())  //1
	fmt.Println(f1()) //5
	fmt.Println(f3()) //1
}

func TestDeferOut(t *testing.T) {
	deferOut()
}

func TestAll(t *testing.T) {
	fRecover()
	fmt.Println("Returned normally from fRecover.")
}
