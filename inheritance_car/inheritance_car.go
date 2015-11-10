package main

import "fmt"

type Engine interface {
	Start()
	Stop()
}

func (c *Car) Start() {
	fmt.Println("Engine Started")
}

func (c *Car) Stop() {
	fmt.Println("Engine Stopped")
}

func (c *Car) numberOfWheels() int {
	return c.Wheel_count
}

type Car struct {
	Wheel_count int
	Engine
}

type Mercedes struct {
	Car
}

func (m *Mercedes) sayHiToMerkel() {
	fmt.Println("Hi")
}

func main() {
	fmt.Println("entered main")
	c := &Car{Wheel_count: 4}
	c.Start()
	fmt.Println(c.numberOfWheels())
	fmt.Println("Lets try the Mercedes")
	m := &Mercedes{Car{Wheel_count: 4}}
	m.Start()
	fmt.Println(m.numberOfWheels())
	m.sayHiToMerkel()
}
