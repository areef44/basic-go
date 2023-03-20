package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func mainasaasdasdsad() {
	person := NewMap("Arif")
	fmt.Println(person)
}
