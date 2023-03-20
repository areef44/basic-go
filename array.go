package main

import "fmt"

func array() {
	var names [2]string

	names[0] = "Muhammad"
	names[1] = "Arif"

	fmt.Println(names[0])
	fmt.Println(names[1])

	var values = [3]int{
		90,
		80,
		70,
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])
	fmt.Println(len(names))
	fmt.Println(len(values))
}
