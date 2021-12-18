package Fundamental

import "fmt"

func Map() {
	var data map[int]string
	//atau
	//var data = map[int]string{}
	data = map[int]string{}

	data[1] = "Fajar"
	data[2] = "Mukhollis"
	// _, isExist := data[1]
	// fmt.Println(isExist)
	delete(data, 1)
	fmt.Println(data)
	//map secara horizontal
	data2 := map[string]string{
		"nama" : "Fajar",
		"NPM" : "4519210034",
	}
	fmt.Println(data2)

	for key, value := range data{
		fmt.Println("Key :", key, "Value :", value)
	}

	//penggunaan make
	mahasiswa := make(map[string]string)
	mahasiswa["nama"] = "Fajar"
	mahasiswa["fakultas"] = "Teknik Informatika"

	fmt.Println(mahasiswa)

	//kombinasi slices and map
	dataMahasiswa := [] map[string]string {
		{"name" : "Fajar", "Kelas" : "Golang"},
		{"name" : "Mukhollis", "Kelas" : "Golang"},
	}
	fmt.Println(dataMahasiswa)

	mhsNew := map[string]string {"name" : "Roy", "Kelas" : "Golang"}

	dataMahasiswa = append(dataMahasiswa, mhsNew)
	fmt.Println(dataMahasiswa)
}