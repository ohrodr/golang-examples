package main

import "fmt"

type Celsius float64

func (c *Celsius) String() string {
	return fmt.Sprintf("%.1f%cC", float64(*c), '\u00b0')
}

func main() {
	c := Celsius(5.44)
	fmt.Println(c.String())
}
