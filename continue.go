package main

import "fmt"

func continuu() {
	for i := 0; i < 10; i++ {
		if i%2 == 1 {
			continue
		}
		fmt.Println("Perulangan ke", i)
	}
}
