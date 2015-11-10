package main

import "fmt"

// calculates fibonacci to 50 using simple array stuff.
func main() {
	// we prepopulate the beginning values
	// this is important bc of the way fib works
	fib := GetFibSlice()
	fmt.Println(fib)
}

func GetFibSlice() func(x int) []int {
	var fib = [50]int{1, 1, 2}
	for i := 3; i < len(fib); i++ {
		fib[i] = fib[i-2] + fib[i-1]
	}
	return func(x int) []int {
		return fib[:x]
	}
}
