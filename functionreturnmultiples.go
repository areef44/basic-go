package main

import "fmt"

func getFullName() (string, string, string) {
	return "Muhammad", "Arif", "Bukittinggi"
}

func maini() {
	firstName, middlename, lastName := getFullName()
	fmt.Println(firstName)
	fmt.Println(middlename)
	fmt.Println(lastName)

	//apabila variabel tidak ingin dipakai tambahkan underscore
	// firstName, _, _ := getFullName()
	// fmt.Println(firstName)
}
