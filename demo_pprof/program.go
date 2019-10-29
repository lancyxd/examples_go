package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
)

/*
go run program.go  --cpuprofile=fabonacci.prof

*/

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			fmt.Println(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
}
