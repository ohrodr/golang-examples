// another example from a book print a rectangle of stars w=20 h=10
package main

import "fmt"

const (
	length int = 10
	height int = 20
)

// this is the main program to print the rectangle
func main() {
	for x := 0; x <= length-1; x++ {
		for y := 0; y <= height-1; y++ {
			switch {
			case x == 0 || y == 0:
				fmt.Print("*")
			case y == height-1 || x == length-1:
				fmt.Print("*")
			default:
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}
