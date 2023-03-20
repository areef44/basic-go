package main

import "fmt"

type HasName interface {
	GetName() string
}

func ucapHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func mainasasaa() {
	var arif Person

	arif.Name = "Arif"

	ucapHello(arif)

	cat := Animal{
		Name: "MengMeng",
	}
	ucapHello(cat)
}
