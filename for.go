package main

import "fmt"

func forange() {

	//init statement                 //post statement
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke", counter)
	}

	//deklarasi array yang mau ditampilkan
	slice := []string{"Muhammad", "Arif", "Joko"}

	//looping expression
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	//shorthand jika ingin nampilin index
	for index, value := range slice {
		fmt.Println("index", index, "=", value)
	}

	//shorthand jika tidak ingin nampilin index ganti indexnya dengan _
	for _, nilai := range slice {
		fmt.Println(nilai)
	}

	//jika menggunakan map
	person := make(map[string]string)
	person["name"] = "Muhammad Arif"
	person["address"] = "Bukittinggi"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}

	//for range bisa digunakan untuk array atau slice akan tetapi bisa juga digunakan untuk map
	//bedanya map gunakan key,kalo slice atau array gunakan index
}
