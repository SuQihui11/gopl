package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

var dilbert Employee

func main() {
	dilbert.Salary -= 5000 // demoted, for writing too few lines of code
	pos := &dilbert.Position
	*pos = "senior" + *pos

	fmt.Printf("%T\n", pos)

	var employeeOfTheMonth *Employee = &dilbert
	employeeOfTheMonth.Position += " (proactive team player)" // (*employeeOfTheMonth).Position += " (proactive team player)"

	fmt.Printf("%T", employeeOfTheMonth.Position)
	fmt.Println(EmployeeByID(dilbert.ManagerID).Position) // "Pointy-haired boss"

	//id := dilbert.ID
	// 如果将EmployeeByID函数的返回值从*Employee指针类型改为Employee值类型，那么更新语句将不能编译通过
	//因为在赋值语句的左边并不确定是一个变量（译注：调用函数返回的是值，并不是一个可取地址的变量）
	//EmployeeByID(id).Salary = 0 // fired for... no real reason
}

func EmployeeByID(id int) Employee { /* ... */ }
