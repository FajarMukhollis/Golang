package Fundamental

import "fmt"

func Array() {
	//Penulisan array
	//Array 1 deminsi
	/*
		var numbers [2]int
		numbers[0] = 1
		numbers[1] = 2
		fmt.Println(numbers)
	*/
	//or
	/*
		var numbers = [2]int {1,2}
		fmt.Println(numbers)
	*/
	//or
	/*
		var numbers = [...]int {1,2,3,4,5,6}
		fmt.println(numbers)
	*/

	//Array 2 dimensi
	// 2 yang pertama -> row (panjang)
	// 2 yang ke dua -> coloum (isi)
	var numbers = [2][2]int{
		{1,2},
		{3,4},
	}
	fmt.Println(numbers)
}
