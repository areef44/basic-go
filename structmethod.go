package main

import "fmt"

type Customer struct {
	Name    string
	Address string
	Age     int
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

func (a Customer) sayHa() {
	fmt.Println("Haaaaaaa from", a.Name)
}

func mainoieuroiuewroiewbdyudbewubd() {
	var customer Customer
	customer.Name = "Muhammad Arif"
	customer.Address = "Bukittinggi"
	customer.Age = 30

	customer.sayHi("Arif")
	customer.sayHa()

	// fmt.Println(customer.Name)
	// fmt.Println(customer.Address)
	// fmt.Println(customer.Age)

	// arif := Customer{
	// 	Name:    "Joko",
	// 	Address: "Cirebon",
	// 	Age:     35,
	// }
	// fmt.Println(arif)

	// budi := Customer{
	// 	"Budi", "Jakarta", 25,
	// }

	// fmt.Println(budi)

}
