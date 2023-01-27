package main

import "fmt"

func main() {
	i := 1;

	// Go doesn't have point arithmetic
	var p *int = nil;

	p = &i // takes the i memory address and gives to p

	// access the value associated with a point using *{variable name} notation
	*p++
	i++

	fmt.Println(p, *p, i)

}