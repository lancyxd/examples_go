package examples_go

import (
	"fmt"
	"testing"
)

func TestEmptyI(t *testing.T) {
	str := "hello"
	i := 5
	strt := struct {
		name string
	}{name: "Lucy"}

	emptyInterface(str)
	emptyInterface(i)
	emptyInterface(strt)
}

func TestAssertI(t *testing.T) {
	str := "hello"
	i := 5

	assertInterface(str)
	assertInterface(i)
}

func TestVowels(t *testing.T) {
	name := MyString("Sam Anderson") //具体的对象
	var v VowelsFinder               //定义接口类型
	v = name

	fmt.Printf("接口类型:%c\n", v.FindVowels())
	fmt.Printf("具体对象:%c\n", name.FindVowels())
}

func TestInShow(t *testing.T) {
	var tValue Test
	fValue := MyFloat(89.7)
	tValue = fValue

	describe(tValue)
	tValue.Tester()
}

func TestEmployee(t *testing.T) {
	oemp1 := OutstandEmployee{1, 3000, 10000}
	oemp2 := OutstandEmployee{2, 3000, 20000}
	cemp1 := CommonEmployee{3, 3000}

	employees := []SalaryCaculator{oemp1, oemp2, cemp1}
	total := totalExpense(employees)
	fmt.Println("TestEmployee total expense:", total)
}

func TestPointer(t *testing.T) {
	var d1 Describer
	p := Person{"lily", 5}
	d1 = p
	d1.Describle()

	p2 := Person{"James", 32}
	d1 = &p2 // 指针类型
	d1.Describle()

	var d2 Describer
	a := Address{"Washington", "USA"}

	//d2 = a  // 不能使用值类型（引发panic）

	d2 = &a
	d2.Describle()
	a.Describle() // 直接使用值类型调用
}

func TestNestedI(t *testing.T) {
	e := Employee{
		firstname:   "Naveen",
		lastname:    "Ramanathan",
		basicpay:    5000,
		pf:          200,
		totalleaves: 30,
		leavestaken: 5,
	}

	var empOp EmployeeOperations = e //声明并初始化
	empOp.DisplaySalary()
	fmt.Printf("left:%d\n", empOp.CalculateLeavesLeft())

	var d3 Describer //空接口
	if d3 == nil {
		fmt.Printf("d3 is nil and 类型 %T 值%v\n", d3, d3)
	}
}
