package main

import "fmt"

func function() {
	var name = "wati"

	if name == "Arif" {
		fmt.Println("Hallo Arif")
	} else if name == "Joko" {
		fmt.Println("Hello Joko")
	} else {
		fmt.Println("Hallo Orang Lain")
	}

	var length = len(name)
	if length > 5 {
		fmt.Println("Terlalu Panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}

}
