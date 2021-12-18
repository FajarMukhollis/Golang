package Fundamental

import "fmt"

func Operators() {

	//num1 := 1
	//num2 := 2

	//Arithmetic Op
	/*
		fmt.Println(num1 + num2)
		fmt.Println(num1 - num2)
		fmt.Println(num1 / num2)
		fmt.Println(num1 * num2)
		fmt.Println(num1 % num2)
		num1++
		num2--
		fmt.Println("Arithmetic ++",num1)
		fmt.Println("Arithmetic --",num2)
	*/

	//Relational
	/*
		result := num1 < num2
		result := num1 > num2
		result := num1 == num2
		result := num1 <= num2
		result := num1 >= num2
		result := num1 != num2

		fmt.Println(result)
	*/

	//Logical Op
	/*
		var left = true
		var right = false
		var leftANDright = left && right
		var leftORright = left || right
		var NOTright = !right
		var NOTleft = !left

		fmt.Println("left AND right : ", leftANDright)
		fmt.Println("left OR right : ", leftORright)
		fmt.Println("NOT right : ", NOTright)
		fmt.Println("NOTleft : ", NOTleft)
	*/

	//Assignment Op
	// = assign a value
	// := deklarasi variable
	// +=
	/*
		num1 += 2
		println(num1)
	*/

	//bitwise Op
	bit1 := 4
	bit2 := 6

	fmt.Println(bit1 & bit2)
	fmt.Println(bit1 | bit2)
	fmt.Println(bit1 ^ bit2)
	fmt.Println(bit1 &^ bit2)
	fmt.Println(bit1 << bit2)
	fmt.Println(bit1 >> bit2)
}
