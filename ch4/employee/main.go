package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID            int
	Name, Address string
	Dob           time.Time
	Position      string
	Salary        int
	ManagerID     int
}

//注意返回值与指针的区别
func EmployeeByID(id int) Employee {

	var tom Employee
	tom = Employee{ID: 100, Name: "Tom", Address: "Beijing"}
	return tom
}

//测试给结构体成员赋值
func testSet() {
	var dibert Employee
	fmt.Println(EmployeeByID(1).Address)
	//EmployeeByID(1).Address = '111'// cannot assign to EmployeeByID(1).Addres
	dibert = EmployeeByID(1)
	dibert.Address = "Shanghai" //可以通过变量赋值
	fmt.Println(dibert.Address)
}

func testDot() {
	var dibert Employee
	//直接修改成员内容
	dibert.Salary += 100
	fmt.Printf("%+v\n", dibert)
	sa := &dibert.Salary
	*sa += 200
	fmt.Printf("%+v", dibert)

}

func main() {
	//var dibert Employee
	//fmt.Printf("%+v", dibert)
	//testSet()
	testDot()
}
