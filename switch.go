package main

import "fmt"

func swit() {

	name := "barjo"

	switch name {
	case "Arif":
		fmt.Println("Hallo Arif")
		fmt.Println("Hallo Arif")
	case "Joko":
		fmt.Println("Hallo Joko")
		fmt.Println("Hallo Joko")
	default:
		fmt.Println("Kenalan Yuk")
		fmt.Println("Kenalan Yuk")
	}

	// switch length := len(name); length > 5 {
	// case true:
	// 	fmt.Println("Nama Terlalu Panjang")
	// case false:
	// 	fmt.Println("Nama Sudah Benar")
	// }

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama Terlalu Panjang")
	case length > 5:
		fmt.Println("Nama lumayan Panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
