// package non-recursive fibonacci
// using closures aka lambdas
package main

import "fmt"

func main() {
	fmt.Println("Fibonacci Table")
	var f_fn = fib()
	for i := 0; i < 10; i++ {
		fmt.Println(i, f_fn())
	}
}

func fib() func() int {
	a := 0
	b := 1
	c := a + b
	return func() int {
		// we basically want to shift the values around
		c, a, b = a, b, a+b
		return c
	}
}
