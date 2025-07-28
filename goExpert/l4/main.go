package main

import "fmt"

type ID int

var userId ID = 20

func main() {

	fmt.Printf("O tipo de userId é %T", userId) // main.ID

	a := 1
	b := 2.334432
	c := true
	d := "opa"

	fmt.Printf("\n%d, %f, %.2f, %t, %s", a, b, b, c, d)
	fmt.Printf("\n%v %v %v %v", a, b, c, d)
}
