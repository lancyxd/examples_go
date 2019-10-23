package examples_go

import (
	"testing"
)

/*
单元测试函数名Test开头，接收一个指针型参数（*testing.T）;func TestXxx(*testing.T)
测试行必须Test开头，Xxx为字符串，第一个X必须大写的[A-Z]的字母，，一般Xxx为被测试方法的函数名。

go test //测试所有
go test -v 或者go test -v  sum_test.go sum.go//测试整个文件
go test -v -test.run TestSub //测试单个函数
go test  -test.run TestforBreak //t.Log不输出,fmt会输出,加-v t.Log输出
go test -v for_select_switch.go for_select_switch_test.go -test.run TestSwitch //指定文件测试
go test -v cs_bench_test.go  cs_unit_test.go

go程序性能分析:go tool pprof
*/

//原函数
func AddTwo(a, b int) int {
	return a + b
}

func SubTwo(a, b int) int {
	return a - b
}

func SumAll(n int) (res int) {
	res = 0
	for i := 1; i <= n; i++ {
		res += i
	}
	return res
}

//单元测试
func TestAdd(t *testing.T) {
	if AddTwo(1, 2) != 3 {
		t.Error("test foo:AddTwo failed")
	} else {
		t.Log("test foo:AddTwo pass")
	}
}

func TestSubTwo(t *testing.T) {
	if SubTwo(2, 1) != 1 {
		t.Error("test foo:SubTwo failed")
	} else {
		t.Log("test foo:SubTwo pass")
	}
}

func TestSum(t *testing.T) {
	if SumAll(5) != 15 {
		t.Error("test foo:SumAll failed")
	} else {
		t.Log("test foo:SumAll pass")
	}
}
