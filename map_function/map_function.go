// Code Challenge write a map function, takes a lambda that multiplies an int by 10
// also takes a list and returns a list of the results applied to the fn passed
package main

import "fmt"

func main() {
	byTen := func(x int) int { return x * 10 }
	res := ListByTen(byTen, []int{1, 2, 3, 4, 5})
	fmt.Println(res)
}

func ListByTen(by func(int) int, these []int) []int {
	fn := by
	ret := make([]int, len(these), cap(these)+1)
	for idx, x := range these {
		ret[idx] = fn(x)
	}
	return ret
}
