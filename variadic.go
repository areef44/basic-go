package main

import "fmt"

//membuat variadic variabel
func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func variadic() {
	total := sumAll(10, 10, 10, 10, 10, 10)
	fmt.Println(total)

	//cara memasukan slice kedalam variabel arguments di variadic function
	slices := []int{10, 10, 10, 10, 10}
	total = sumAll(slices...)
	fmt.Println(total)
}
