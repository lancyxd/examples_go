//go build -buildmode=plugin -o aplugin.so aplugin.go

package main

import (
	"fmt"
)

func Add(x, y int) int {
	fmt.Println("begin excu Add v1")
	return x + y
}

func Sub(x, y int) int {
	fmt.Println("begin excu Sub v1")
	return x - y
}
