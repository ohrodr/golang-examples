// another example populating an array and printing the result
// also demonstrates indexing with range as well as _
package main

import "fmt"

func main() {
	var a [16]int
	for index, _ := range a {
		a[index] = index
	}
	fmt.Println(a)
}
