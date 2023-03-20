package main

import "fmt"

func variable() {
	var name string

	name = "Muhammad Arif"
	fmt.Println(name)

	name = "Joko Widodo"
	fmt.Println(name)

	//atau

	var nama = "Muhammad Arif"
	fmt.Println(nama)

	nama = "Mas Dodo"
	fmt.Println(nama)

	var age = 38
	fmt.Println(age)

	country := "indonesia"
	fmt.Println(country)

	var (
		namaAwal  = "Muhammad"
		namaAkhir = "Arif"
	)

	fmt.Println(namaAwal)
	fmt.Println(namaAkhir)

}
