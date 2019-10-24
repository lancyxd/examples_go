/*
interface可以定义一组方法，但是这些不需要实现
interface是一种类型，确切的说，是指针类型

*/
package examples_go

import (
	"fmt"
)

//1)空接口
func emptyInterface(i interface{}) {
	fmt.Printf("type: %T,value: %v\n", i, i)
}

//2)类型断言
func assertInterface(i interface{}) {
	v, ok := i.(int) //类型断言，防止产生panic
	fmt.Println("assertInterface:", v, ok)

	switch i.(type) {
	case string:
		fmt.Printf("type string:%s\n", i.(string))
	case int:
		fmt.Printf("type int:%d\n", i.(int))
	default:
		fmt.Printf("unkonwn type")
	}
}

//3)创建和实现接口
type VowelsFinder interface {
	FindVowels() []rune
}

type MyString string //具体的对象及行为
func (s MyString) FindVowels() []rune {
	var vowels []rune
	for _, rune := range s {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
}

//4)接口内部表现
type Test interface {
	Tester()
}

type MyFloat float64

func (m MyFloat) Tester() {
	fmt.Println("Tester:", m)
}
func describe(t Test) {
	fmt.Printf("Interface 类型 %T ,  值： %v\n", t, t)
}

//接口实际用途
type SalaryCaculator interface {
	CaculateSalary() int
}

type CommonEmployee struct {
	empId    int
	basicpay int
}

type OutstandEmployee struct {
	empId    int
	basicpay int
	jj       int //奖金
}

func (c CommonEmployee) CaculateSalary() int {
	return c.basicpay
}
func (o OutstandEmployee) CaculateSalary() int {
	return o.basicpay + o.jj

}
func totalExpense(s []SalaryCaculator) int {
	expense := 0
	for _, v := range s {
		expense = expense + v.CaculateSalary()
	}

	return expense
}

//5)指针和值类型
type Describer interface {
	Describle()
}

type Person struct {
	name string
	age  int
}

func (p Person) Describle() { //使用值接收器实现
	fmt.Printf("%s is %d years old\n", p.name, p.age)
}

type Address struct {
	state   string
	country string
}

func (a *Address) Describle() { //使用指针接收器实现
	fmt.Printf("State %s Country %s\n", a.state, a.country)
}

//6)接口嵌套
type SalaryCalculator interface {
	DisplaySalary()
}
type LeaveCalculator interface {
	CalculateLeavesLeft() int
}

type EmployeeOperations interface {
	SalaryCalculator
	LeaveCalculator
}

type Employee struct {
	firstname   string
	lastname    string
	basicpay    int
	pf          int
	totalleaves int
	leavestaken int
}

func (e Employee) DisplaySalary() {
	fmt.Printf("%s %s has salary $%d\n", e.firstname, e.lastname, (e.basicpay + e.pf))
}

func (e Employee) CalculateLeavesLeft() int {
	return e.totalleaves - e.leavestaken
}
