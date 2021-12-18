package Fundamental

import "fmt"

func Pointer() {
	var angka int = 1
	// var angka1 *int
	// angka1 = &angka
	// fmt.Println(&angka)
	// fmt.Println(*angka1)
	fmt.Println("Angka sebelum berubah : ", angka)
	changeNumber(&angka, 10)
	fmt.Println("Angka setelah berubah : ", angka)
}

func changeNumber(number *int, value int) {
	*number = value
	fmt.Println(*number)
}