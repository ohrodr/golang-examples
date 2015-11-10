// simple app proves arrays are copies within functions
package main

import "fmt"

func main() {
	var a [5]int
	showArray(a, a)
	fmt.Println("Notice left side doenst change")
	showArrayCopyRef(&a, &a)
}

// given a reference to an array
// run the dupe function
func showArrayCopyRef(vals *[5]int, valz *[5]int) {
	showArrayCopy(*vals, *valz)
}

// dupearray function shows the array is a copy
func showArrayCopy(vals [5]int, valz [5]int) {
	for i := 0; i < len(vals); i++ {
		fmt.Println(i, vals, valz)
		valz[i] = 1
		fmt.Println(i, vals, valz)
	}

}
