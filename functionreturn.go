package main

import "fmt"

//jika ingin mengembalikan data tambahkan tipe data yang akan dikembalikan
func getHello(name string) string {
	// result := "Hello"
	// return (result)
	if name == "" {
		return "Hello Bro"
	} else {
		return "Hello " + name
	}

}

func oke() {
	result := getHello("Arif")
	fmt.Println(result)

	fmt.Println(getHello(""))
}
