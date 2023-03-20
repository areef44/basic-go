package main

import "fmt"

type Alamat struct {
	City, Province, Country string
}

//menambahkan pointer function bisa dengan menambahkan * di struct alamat
func ChangeCountryToWakanda(alamat *Alamat) {
	alamat.Country = "Indonesia"
}

func mainpipoiopi() {

	var alamat = Alamat{
		City:     "Palembang",
		Province: "Sumatera Selatan",
		Country:  "",
	}

	//dan untuk pemanggilan di functionnya tambahkan & di parameter structnya
	var alamatPointer *Alamat = &alamat
	ChangeCountryToWakanda(alamatPointer)
	fmt.Println(alamatPointer)
}

/**
ketika ingin membuat data struct yang lumayan besar ketika ingin membuat parameter alangkah
baiknya dijadikan pointer karena tiap manggil function kalo tidak dijadikan pointer maka
datanya akan selalu di duplicate di memory,dan pastinya memory nya akan membengkak
*/
