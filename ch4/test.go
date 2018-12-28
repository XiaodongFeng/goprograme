package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID int
	Name, Address string
	DoB time.Time
	Position string
	Salary int
	ManagerID int
}

func EmployeeByID(id int) *Employee {
	e := &Employee{ID: id}
	return e
}


func main() {
	a := [...]string{1:"1", 2:"2", 5:"5", 4:"4"}
	fmt.Println(a, len(a))
	b := make(map[string]int)
	b["a"] = 1
	c := map[string]int{}
	c["a"] = 1
	//var d map[string]int
	//d["a"] = 1
	var dilbert Employee
	fmt.Println(EmployeeByID(dilbert.ManagerID).Position) // "Pointy-haired boss"
	id := dilbert.ID
	EmployeeByID(id).Salary = 0
}
