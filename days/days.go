package main

import "fmt"

type Day int

var names = []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

const (
	SN = iota
	MO
	TU
	WE
	TH
	FR
	SA
)

func (d *Day) String() string {
	return fmt.Sprintf(names[int(*d)])
}

func main() {
	c := Day(MO)
	fmt.Println(c.String())
}
