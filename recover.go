package main

import "fmt"

func endAppTwo() {
	message := recover()
	if message != nil {
		fmt.Println("Error Dengan Pesan:", message)
	}
	fmt.Println("Aplikasi Selesai")
}

func runAppTwo(error bool) {
	defer endAppTwo()
	if error {
		panic("ERROR")
	}
	fmt.Println("Aplikasi Berjalan")
}

//recover digunakan menangkap string error
func mainees() {
	runAppTwo(true)
	fmt.Println("Arif")
}
