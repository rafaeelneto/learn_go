package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 3.23
	y := 12

	fmt.Println(x / float64(y))
	
	grade := 6.9
	finalGrade := int(grade)
	fmt.Println(finalGrade)

	// WARNING
	fmt.Println("Teste", string(123))

	// int to string
	fmt.Println("Teste", strconv.Itoa(123))

	// string to int
	num, _ := strconv.Atoi("6468")
	fmt.Println(num)

	b, _ := strconv.ParseBool("true")
	if b {
		fmt.Println("Verdade")
	}
}