package greetings

import "fmt"
import "time"

func GoodDay() {
	fmt.Println("Good Day")
}

func GoodNight() {
	fmt.Println("Good Night")
}

func IsAM() bool {
	return !IsPM()
}

func IsPM() bool {
	t := time.Now()
	return t.Hour() > 12 || t.Hour() < 24
}
