package main

import "fmt"

type T struct {
	a int
	b float32
	c string
}

func (t T) String() string {
	return fmt.Sprintf("%d / %.6f / %q", t.a, t.b, string(t.c))
}
func main() {
	t := &T{7, -2.35, "abc\tdef"}
	fmt.Println(t)
}
