package main

import (
	"fmt"
	"golang-vp/Fundamental"
)

//GLOBAL VARIABLE
var NAMA string = "Fajar Mukhollis"
var lahir string = "10 September 2000"

func main() {
	/*
		//deklarasi variabel
		//cara ke-1
		//string
		var uji string = "Test"
		//cara ke-2
		uji2 := "test lagi"
		huruf := 'A'
		//boolean
		isActive := true
		//int8 (-128 sampai 127)
		var sampleint8 int8 = 126
		//int16 (-32768 sampai 32767)
		//int32 (-2147483648s sampai 2147483647)
		//int64 (-9223372036854775808 sampai 9223372036854775807)
		//rune (sama dengan int32)
		//untuk lebih lengkap silakan masuk ke website : https://dasarpemrogramangolang.novalagung.com/A-tipe-data.html
		// zero value -> tidak memilii value, nilainya default
		var number1 int
		var value bool

		const phi = 3.14

		//untuk mencetak
		fmt.Println(uji, "GOLANG")
		fmt.Println(uji2)
		fmt.Println(huruf)
		fmt.Println(isActive)
		fmt.Println(sampleint8)
		fmt.Println(number1)
		fmt.Println(value)
		fmt.Println(phi)
		Greet()

		Fundamental.Operators()
		Fundamental.Condition()
		Fundamental.Loop()
		Fundamental.Array()
		Fundamental.Slices()
		Fundamental.Map()
		Fundamental.Function("Dunia ", "Fajar ", 4519210034)
		// result := Fundamental.Sum(1,2)
		// fmt.Println(result)
		Fundamental.SetNames("Fajar", "Mukhollis")
		Fundamental.Variblefunc()
		Fundamental.Pointer()
		Fundamental.Input()
		Fundamental.Struct()
		//Method
		// rect := Fundamental.Rectangle {
		// 	Width: 10,
		// 	Lenght: 5,
		// }
		// rect.GetWidht()
		//instance rectangle
		rect := Fundamental.Rectangle{}
		rect.SetWidht(10)
		get := rect.GetWidht()
		fmt.Println(get)
	*/
	//interface
	var shape Fundamental.Shape

	shape = Fundamental.Rectangles{Width: 10.0, Length: 2.0}
	fmt.Println(shape.GetArea())

	shape = &Fundamental.Circle{}
	shape.(*Fundamental.Circle).SetRadius(10.0)
	fmt.Println(shape.GetArea())

}

func Greet() {
	fmt.Println("Hallo", NAMA)
	fmt.Println("Tanggal Lahir", lahir)
}

// func Calculate (x int) (result int) {
// 	result = x + 2
// 	return result
// }
