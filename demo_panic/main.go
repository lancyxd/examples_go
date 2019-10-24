package main

import (
	"common/logging"
	"examples_go/demo_panic/pfile"
	"fmt"
)

// syscall.NewLazyDLL(kernel32dll)需在windows下运行该文件
func main() {
	fmt.Println("main begin")

	suffix := "filetest"
	file, err := pfile.DumpPanic(suffix)
	if err != nil {
		logging.Error("main DumpPanic err=%s", err.Error())
	}

	defer func() {
		if err := pfile.ReviewDumpPanic(file); err != nil {
			logging.Error("defer func() review dump panic error: %s", err.Error())
		}
	}()

	fmt.Println("main end")
}
