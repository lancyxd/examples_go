/*
go test -bench=. -cpuprofile=cpu.prof 产生*.prof文件
*/

package main

import (
	"testing"
)

const url = "https://github.com/EDDYCJY"

//单元测试
func TestAdd(t *testing.T) {
	s := Add(url)
	if s == "" {
		t.Errorf("Test.Add error!")
	}
}

//性能测试
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(url)
	}
}
