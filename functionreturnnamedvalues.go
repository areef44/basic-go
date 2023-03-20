package main

import "fmt"

func getFullNameTwo() (firstName string, middleName string, lastName string) {
	firstName = "Muhammad"
	middleName = "Arif"
	lastName = "Bukittinggi"
	return
}

func mainu() {
	a, b, c := getFullNameTwo()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
