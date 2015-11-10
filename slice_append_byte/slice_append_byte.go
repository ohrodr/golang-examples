// example challenge
// write a fn to copy a []byte to an []int
package main

import "fmt"

func main() {
	v := make([]int, 10, 50)
	v[1] = 1
	b := []byte{10, 20, 30}
	fmt.Println(v)
	fmt.Println("length:", len(v))
	fmt.Println("cap:", cap(v))
	fmt.Println(b)
	c, err := Append(v, b)
	if err {
		fmt.Println("Byte array bigger than slice capacity")
	}
	fmt.Println(c)
}

// copys byte_array to slice of ints returning err if it wont fit
func Append(slice []int, byte_array []byte) ([]int, bool) {
	if len(byte_array) > cap(slice)-len(slice) {
		return nil, true
	}
	s := make([]int, len(slice)+len(byte_array))
	copy(s, slice)
	for i := 0; i < len(byte_array); i++ {
		s[len(slice)+int(i)] = int(byte_array[i])
	}
	return s, false
}
