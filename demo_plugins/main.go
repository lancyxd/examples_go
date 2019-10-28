//使用so文件
package main

import (
	"fmt"
	"plugin"
)

func main() {
	//加载
	p, _ := plugin.Open("plugins/aplugin.so")

	//调用,使用
	add, _ := p.Lookup("Add")
	addRes := add.(func(int, int) int)(11, 2)
	fmt.Printf("addRes=%d\n", addRes)

	sub, _ := p.Lookup("Sub")
	subRes := sub.(func(int, int) int)(11, 2)
	fmt.Printf("subRes=%d\n", subRes)

}
