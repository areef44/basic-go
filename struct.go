package main

import "fmt"

type Customers struct {
	Name    string
	Address string
	Age     int
}

func mainsss() {
	var customers Customer
	customers.Name = "Muhammad Arif"
	customers.Address = "Bukittinggi"
	customers.Age = 30

	fmt.Println(customers.Name)
	fmt.Println(customers.Address)
	fmt.Println(customers.Age)

	arif := Customer{
		Name:    "Joko",
		Address: "Cirebon",
		Age:     35,
	}
	fmt.Println(arif)

	budi := Customer{
		"Budi", "Jakarta", 25,
	}

	fmt.Println(budi)

}
