package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile(`e([a-z])o`)

	fmt.Println(regex.MatchString("ako"))
	fmt.Println(regex.MatchString("ato"))
	fmt.Println(regex.MatchString("aDo"))

	search := regex.FindAllString("ero eno edo eki", -1)
	fmt.Println(search)
}
