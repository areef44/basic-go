package main

import (
	"fmt"
)

//this is a comment

func hello() {
	fmt.Println("Hello World")

	fmt.Println("Satu = ", 1)
	fmt.Println("Dua = ", 2)
	fmt.Println("Tiga Koma Lima", 3.5)

	fmt.Println("Benar =", true)
	fmt.Println("Salah =", false)

	fmt.Println("String")
	fmt.Println("Tipe Data String")
	fmt.Println("Ini String")

	//ngitung panjang string
	fmt.Println(len("ARIF"))

	//milih string berdasarkan index
	fmt.Println("ARIF"[0])
	fmt.Println("arif"[0])
}
