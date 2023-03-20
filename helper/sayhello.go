package helper

import "fmt"

//ini tidak bisa diakses di paket lain
var version = "1.0.0"

//ini bisa diakses di paket lain
var Application = "Belajar Golang"

//ini bisa diakses di paket lain
func BilangHello(name string) {
	fmt.Println("hello", name)
}

//ini tidak bisa diakses di paket lain
func bilangDadah(name string) {
	fmt.Println("Goodbye", name)
}
