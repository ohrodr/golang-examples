package main

import "fmt"
import "github.com/ohrodr/golang-examples/greetings"

func main() {
	fmt.Println("entering main")
	greetings.GoodDay()
	greetings.GoodNight()
	fmt.Println("Is it morning yet my son asks?", greetings.IsPM())
	fmt.Println("Is it morning yet my son asks?", greetings.IsAM())
}
