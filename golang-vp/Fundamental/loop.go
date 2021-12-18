package Fundamental

import "fmt"

func Loop() {
	/*
		for i := 1; i <= 10; i++ {
			fmt.Println("Loop Ke-", i)
		}
	*/
	/*
		for i := 1; i <= 10; i++ {
			if i%2 == 0 {
				fmt.Println(i)
			}
		}
	*/

	// for while
	/*
		num1 := 1
		for num1 < 10 {
			fmt.Println(num1)
			num1++
		}
	*/

	//for
	/*
		num1 := 1
		for num1 == 1 {
			fmt.Println("this command will be execute again and again forever")
			if num1 == 1 {
				fmt.Println("Kondisi terpenuhi")
				break
			}
		}
	*/
	/*
		num1 := 10
		for num1 < 18 {
			if num1 == 15 {
				num1 = num1 + 1
				continue
			}
			if num1 == 16 {
				break
			}
			fmt.Println("Value", num1)
			num1++
		}
	*/

	//Defer
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")

}
