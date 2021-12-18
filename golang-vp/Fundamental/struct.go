package Fundamental

import "fmt"

type Car struct {
	//Property
	name  string
	color string
	fuel  int
}

func Struct() {
	var GreetCorolla = Car{}
	GreetCorolla.name = "Tesla"
	GreetCorolla.color = "Navy"
	GreetCorolla.fuel = 30
	fmt.Println(GreetCorolla)

	//menggunakan pointer
	var allNew *Car = &GreetCorolla
	fmt.Println(GreetCorolla, *allNew)
	allNew.color = "Dark Blue"
	fmt.Println(GreetCorolla, *allNew)
}
