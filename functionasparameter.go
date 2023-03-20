package main

import (
	"fmt"
)

// type declarations untuk dipanggil di function sayhellowith filter
type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func mainx() {
	sayHelloWithFilter("Arif", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Anjing", filter)
}
