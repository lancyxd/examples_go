package examples_go

import (
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
	"testing"
)

/*
基准测试以Benchmark开头，接收一个指针型参数（*testing.B）;性能测试func BenchmarkXxx(*testing.B)

go test -bench=".*" //运行所有基准测试
go test ./bench_test.go -bench=".*" //对某个go文件进行benchmark测试
go test -bench="BenchmarkSequential" //-bench可以指定函数名，支持正则
go test -bench="BenchmarkSequential" -benchmem  //-benchmem 基准测试中包含内存分配信息
go test -bench="BenchmarkSequential" -benchmem  -benchtime 10ms//-benchtime t 用来间接控制基准测试函数的操作次数 t时间 1s
go test -bench="BenchmarkConcurrent" -benchmem -cpu=1,4,16 //-cpu 自定义测试运行次数，并在测试运行时间改变go语言的最大并发处理数的标记

b.SetBytes(1024) MB/s 每秒被处理的字节数量
ns/op 平均耗时
B/op 每次操作分配的字节平均数
allocs/op 每次操作分配内存的次数
go test -v -run="none" -bench=.    不允许单元测试，运行所有的基准测试，

GOGC=off go test -cpu 1 -run none -bench . -benchtime 3s
GOGC=off go test -cpu 8 -run none -bench . -benchtime 3s
*/

//基准测试函数
func BenchmarkSequential(b *testing.B) {
	numbers := []int{1, 3, 5, 7, 9, 11}
	b.N = 100000000
	for i := 0; i < b.N; i++ {
		Add(numbers)
	}
}

func BenchmarkConcurrent(b *testing.B) {
	numbers := []int{1, 3, 5, 7, 9, 11}
	b.N = 100000000
	for i := 0; i < b.N; i++ {
		AddConcurrent(runtime.NumCPU(), numbers)
	}
}

//性能测试
func BenchmarkAdd(b *testing.B) {
	// 如果需要初始化，比较耗时的操作可以这样：
	// b.StopTimer()
	// .... 一堆操作
	// b.StartTimer()
	b.N = 100000000
	for i := 0; i < b.N; i++ {
		AddTwo(1, 2)
	}
}

//原函数
func Add(nums []int) int {
	var v int
	for _, i := range nums {
		v += i
	}
	return v
} //将输入切片分解，然后同时处理他们，然后把小切片的结果相加

func AddConcurrent(goroutines int, nums []int) int {
	var v int64
	totalNumbers := len(nums)
	lastgoroutions := goroutines - 1
	stride := totalNumbers / goroutines

	var wg sync.WaitGroup
	wg.Add(goroutines)
	for g := 0; g < goroutines; g++ {
		go func(g int) {
			start := g * stride
			end := start + stride
			if g == lastgoroutions {
				end = totalNumbers
			}

			var lv int
			for _, n := range nums[start:end] {
				lv += n
			}
			atomic.AddInt64(&v, int64(lv))
			wg.Done()
		}(g)
	}
	wg.Wait()
	return int(v)
}

//性能测试
func BenchmarkItoaStrconv(b *testing.B) {
	for n := 0; n < b.N; n++ {
		result = strconv.Itoa(n)
	}
}

func BenchmarkItoaFmt(b *testing.T) {
	for n := 0; n < b.N; n++ {
		result = strconv.Itoa(n)
	}

}
