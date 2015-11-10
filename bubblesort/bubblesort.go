// Challenge implement bubblesort which sorts a slice of []int
package main

import "fmt"

func main() {
	v := []int{5, 1, 4, 2, 8}
	fmt.Println(v)
	BubbleSort(v)
	fmt.Println(v)
}

// Bubblesort Side-effect sorts via bubble sort
func BubbleSort(val []int) []int {
	found := false
	n := len(val) - 1
	for i := 0; i < n; i++ {
		if val[i] > val[i+1] {
			found = true
			val[i], val[i+1] = val[i+1], val[i]
		}
	}
	if found {
		return BubbleSort(val)
	} else {
		return val
	}
}
