package main

import "fmt"

func main() {
	fmt.Println(sum(3, 4, 23, 5, 5, 6, 6, 4234, 13, 1231, 231231, 23, 123))

	total := func() int {
		return sum(3, 4, 23, 5, 5, 6, 6, 4234, 13, 1231, 231231, 23, 123)
	}()

	fmt.Println(total)
}

func sum(numeros ...int) int {
	total := 0

	for _, numero := range numeros {
		total = total + numero
	}

	return total
}
