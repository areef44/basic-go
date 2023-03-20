package database

import "fmt"

var connection string

// function yang akan dieksekusi pertama kali
func init() {
	fmt.Println("Init dipanggil")
	connection = "MySQL"
}

func GetDatabase() string {
	return connection
}
