package main

import "fmt"

func endApp() {
	fmt.Println("Aplikasi Berakhir")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("ERROR")
	}
	fmt.Println("Aplikasi Berjalan")
}
func mainsuit() {
	runApp(true)
}
