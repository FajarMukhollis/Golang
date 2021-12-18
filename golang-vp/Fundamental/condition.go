package Fundamental

import "fmt"

func Condition() {
	// angka1 := 2
	// angka2 := 2
	// angka3 := 3
	//single statemen
	/*
		if angka1 > angka2 {
			fmt.Println("Lebih Kecil")
		}
	*/

	// if else statemen
	/*
		if angka1 > angka2 {
			fmt.Println("Lebil kecil")
		} else {
			fmt.Println("Lebih besar")
		}
	*/

	// else if
	/*
		if angka1 < angka2 {
			fmt.Println("Lebil kecil")
		} else if angka1 == angka2 {
			fmt.Println("Angka Sama besar")
		} else {
			fmt.Println("Lebih besar")
		}
	*/

	//nested if
	/*
		if angka1 == angka2 {
			if angka1 != angka3 {
				fmt.Println("Kondisi Terpenuhi")
			}
		}
	*/

	//switch
	day := "senin"
	switch day {
	case "senin":
		fmt.Println("Hari Senin")
	case "selasa":
		fmt.Println("Hari Senin")
	default:
		fmt.Println("Waktunya Libur")
	}
}
