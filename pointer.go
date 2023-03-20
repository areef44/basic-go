package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func mainasdsadasdsadsad() {
	address1 := Address{"Bukittinggi", "Sumatera Barat", "Indonesia"}
	//cara menambahkan pointer di golang untuk entity saja menggunakan tanda & di tujuan yang mau dipointer
	//cara kedua var address2 *Address = &address1
	address2 := &address1
	address3 := &address1

	address2.City = "Bandung"

	//cara menambahkan pointer di golang untuk keseluruhan data menggunakan tanda * di data yang mau dipointer
	*address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	//menambahkan pointer dan isi data kosong menggunakan new
	address4 := new(Address)
	address4.City = "Medan"
	fmt.Println(address4)
}
