package main

import "fmt"

func mainua() {
	name := "arif"
	counter := 0
	increment := func() {
		name := "Joko"
		fmt.Println("Increment")
		counter++
		fmt.Println(name)
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)

	//closure = sebuah block scope bisa berinteraksi dengan block scope yg lain
}
