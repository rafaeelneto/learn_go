package main

import "fmt"

func main() {
	// each type has a default zero value
	var a int
	var b float64
	var c bool
	var d string
	var e *int

	fmt.Printf("%v %v %v %q %v", a, b, c, d, e)

}