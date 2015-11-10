package main

import "fmt"

func main() {
	foo := "foobarbaz"
	res1, res2 := Split(foo, 3)
	fmt.Println(res1, res2)
	fmt.Println(res1 == foo)
}

func Split(this string, position int) (string, string) {
	return this[:position], this[position:]
}
