package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (value UserSlice) Len() int {
	return len(value)
}

// sorrt dari kecil ke besar
func (value UserSlice) Less(i, j int) bool {
	return value[i].Age < value[j].Age
}

//sort dari besar ke kecil
// func (value UserSlice) Less(i, j int) bool {
// 	return value[i].Age < value[j].Age
// }

func (value UserSlice) Swap(i, j int) {
	value[i], value[j] = value[j], value[i]
}

func main() {
	users := []User{
		{"Arif", 35},
		{"Joko", 30},
		{"Budi", 25},
		{"Barjo", 31},
		{"Tarno", 34},
	}

	fmt.Println(users)

	sort.Sort(UserSlice(users))

	fmt.Println(users)
}
