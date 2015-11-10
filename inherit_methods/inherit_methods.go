package main

import "fmt"

type Base struct {
	id int
	IDManager
}

type IDManager interface {
	Id()
	SetId()
}

func (b *Base) Id() int {
	return b.id
}

func (b *Base) SetId(val int) {
	b.id = val
}

type Person struct {
	Base
	FirstName string
	LastName  string
}

type Employee struct {
	Person
	salary int
}

func main() {
	e := &Employee{Person{FirstName: "ro", LastName: "dr"}, 1000000}
	fmt.Println("FN:", e.FirstName, "SN:", e.LastName)
	fmt.Println("ID:", e.Id())
	fmt.Println("Salary:", e.salary)
}
