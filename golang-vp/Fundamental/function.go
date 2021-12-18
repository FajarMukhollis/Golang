package Fundamental

import "fmt"

func Function(value, name string, number int) {
	fmt.Println("Hallo " + value + name, number)
}

//mengembalikan nilai dan fungsi return
func Sum(angka1, angka2 int) int {
	return angka1 + angka2
}

//fungsi variadic menggunakan keyword ...
func SetNames(name string, names ... string) {
	fmt.Println(name, names)
}

//function closure
func Variblefunc() {
	greeting := func () {
		fmt.Println("Hello")
	}
	greeting()
}