package study1

import (
	"fmt"
	"math"
)

// 定义接口
type Shape interface {
	Area() float64
	Perimeter() float64
	PrintName() string
}

// 定义矩形结构体
type Rectangle struct {
	Width  float64
	Height float64
	Name   string
}

// 实现Shape 接口的方法
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}

func (r Rectangle) PrintName() string {
	return r.Name
}

// 定义圆形结构体
type Circle struct {
	Redius float64
	Name   string
}

// 实现Shape接口的方法
func (c Circle) Area() float64 {
	return math.Pi * c.Redius * c.Redius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Redius
}

func (c Circle) PrintName() string {
	return c.Name
}

// 题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。

// 定义Person 员工结构体
type Person struct {
	Name string
	Age  int
}

// 定义 Employee 结构体
type Employee struct {
	Person     Person
	EmployeeId int
}

func (emp Employee) PrintInfo() {
	fmt.Printf("员工姓名为：%s, 年龄为: %d, 员工ID为: %d", emp.Person.Name, emp.Person.Age, emp.EmployeeId)
}
