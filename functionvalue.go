package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye " + name
}

func mains() {
	goodBye := getGoodBye

	result := goodBye("Arif")
	fmt.Println(result)
}
