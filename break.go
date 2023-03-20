package main

import "fmt"

func breake() {

	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}

		fmt.Println("Perulangan ke ", i)
		// if i == 5 {
		// 	break
		// }
	}
}
