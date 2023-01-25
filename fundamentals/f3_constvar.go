package main

import "fmt"

func main() {
	// Verbose way
	const PI float64 = 3.123
	var raio = 32.31

	// reduced way of declaring variables
	area := 2121.1

	multiplyBy := 323

	// unused variables generate a compile error
	fmt.Println(raio, area, multiplyBy)

	const (
		a = 1
		b = 2
	)

	var (
		c = 1
		d = 2
	)

	fmt.Println(c, d)


	var e, f bool = true, false;

	fmt.Println(e, f)

	g, h, i := "string qualquer", false, 3.24
	fmt.Println(g, h, i)
	
}