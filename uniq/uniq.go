// challenge write a program that reverse a string, so "Google" is printed as "elgooG"
package main

import "fmt"

func main() {
	str := []byte("Google")
	str2 := Uniq(str)
	fmt.Println(str, str2)
}

// given a slice of bytes, return the uniq vals
// where uniq is if diff than preceeding.
func Uniq(str []byte) []byte {
	ret := make([]byte, 0, cap(str))
	for i := 0; i < len(str); i++ {
		if i == 0 {
			ret = append(ret, str[0])
			continue
		} else if str[i] != str[i-1] {
			ret = append(ret, str[i])
		}
	}
	return ret
}
