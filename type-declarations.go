package main

import (
	"fmt"
)

func typedeclarations() {
	type NoKtp string
	type Married bool

	var noNoKtp NoKtp = "1665487978789"
	var marriedStatus Married = true
	fmt.Println(noNoKtp)
	fmt.Println(marriedStatus)
}
