package main

import "fmt"

type Man struct {
	Name string
}

//membuat pointer di method dengan menambahkan * di struct nya
func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func mainmsnadbasjd() {
	orang := Man{"Arif"}
	orang.Married()

	fmt.Println(orang.Name)
}

//cara ini lebih hemat memory karena tidak menduplicate datanya
