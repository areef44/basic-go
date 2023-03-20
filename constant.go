package main

import "fmt"

func constant() {

	const (
		firstName string = "Muhammad"
		lastName         = "Arif"
		value            = 1000
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(value)
}
