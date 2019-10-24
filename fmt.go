package examples_go

import (
	"fmt"
	"math"
)

type Point struct {
	X int
	Y int
}

func main() {
	var e interface{} = 2.71828
	fmt.Printf("%v(%T)\n", e, e)
	fmt.Printf("%10d\n", 353)
	fmt.Printf("%*d\n", 10, 353)

	nums := []int{12, 237, 3878, 3}
	size := alignSize(nums)
	for i, n := range nums {
		fmt.Printf("%02d %*d\n", i, size, n)
	}

	fmt.Printf("The price of %[1]s was $%[2]d. $%[2]d! imagine that.\n", "carrot", 23)
	p := &Point{1, 2}
	fmt.Printf("%v %+v %#v\n", p, p, p) //%v will print a Go value, it can be modified with + prefix to print field names in a struct and with # prefix to print field names and type.

}

func alignSize(nums []int) int {
	size := 0
	for _, n := range nums {
		if s := int(math.Log10(float64(n))) + 1; s > size {
			size = s
		}
	}
	return size
}
