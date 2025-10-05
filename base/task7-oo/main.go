package main

import "fmt"

func main() {
	s1 := Rectangle{5, 10}
	fmt.Printf("Rectangle#Area:%.2f, Perimeter:%.2f\n", s1.Area(), s1.Perimeter())

	s2 := Circle{5}
	fmt.Printf("Circle#Area:%.2f, Perimeter:%.2f\n", s2.Area(), s2.Perimeter())

	println("-----------------------------")

	emplee := Employee{
		EmployeeID: 1001,
		Person: Person{
			Name: "张三",
			Age: 18,
		},
	}

	println(emplee.GetInfo())
}

// 定义形状的接口
type Shape interface {
	// 计算面积
	Area() float64
	// 计算周长
	Perimeter() float64
}

type Rectangle struct {
	width, height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.radius
}




type Person struct {
	Name string
	Age int
}


type Employee struct {
	Person
	EmployeeID int
}

func (e Employee) GetInfo() string {
	return fmt.Sprintf("Name:%s, Age:%d, EmployeeID:%d", e.Name, e.Age, e.EmployeeID)
}
