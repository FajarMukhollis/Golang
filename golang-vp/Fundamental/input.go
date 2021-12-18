package Fundamental

import (
	"bufio" // -> untuk mengimplementasikan input dan output
	"fmt"
	"os" // ->
)

func Input() {
	var angka int
	fmt.Print("Masukkan angka : ")
	fmt.Scan(&angka)
	fmt.Println("Angka yang dimasukkan : ", angka)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Masukkan kalimat : ")
	scanner.Scan()
	fmt.Println("Kalimat yang dimasukkan : '" + scanner.Text() + "'")
}
