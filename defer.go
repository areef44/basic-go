package main

import "fmt"

func logging() {
	fmt.Println("Selesai Memanggil Function")
}

func runApplication(value int) {
	defer logging()
	fmt.Println("Run Application")
	result := 10 / value
	fmt.Println("Result", result)
}

func mainset() {
	runApplication(10)
}
