package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()

	data.PushBack("Muhammad")
	data.PushBack("Arif")
	data.PushBack("Bukittinggi")
	data.PushFront("Budi")

	fmt.Println(data.Front().Prev())
	fmt.Println(data.Back().Next())

	// dari depan ke belakang
	// for element := data.Front(); element != nil; element = element.Next() {
	// 	fmt.Println(element.Value)
	// }

	// dari belakang ke depan
	for element := data.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
	}
}
