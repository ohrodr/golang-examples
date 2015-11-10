package main

import "fmt"

// calculates fibonacci to 50 using simple array stuff.
func main() {
	// we prepopulate the beginning values
	// this is important bc of the way fib works
	var fib = [50]int{1, 1, 2}
	for i := 3; i < len(fib); i++ {
		fib[i] = fib[i-2] + fib[i-1]
	}
	fmt.Println(fib)
}
