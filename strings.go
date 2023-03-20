package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Muhammad Arif", "Arif"))
	fmt.Println(strings.Contains("Muhammad Arif", "Joko"))

	fmt.Println(strings.Split("Muhammad Arif", " "))

	fmt.Println(strings.ToLower("Muhammad Arif"))
	fmt.Println(strings.ToUpper("Muhammad Arif"))
	fmt.Println(strings.ToTitle("muhammad arif"))

	fmt.Println(strings.Trim("   Muhammad Arif  ", " "))

	fmt.Println(strings.ReplaceAll("   Muhammad Arif  ", "Arif", "Joko"))
}
