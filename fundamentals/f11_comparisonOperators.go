package main

import (
	"fmt"
	"time"
)

func main() {

	// case sensitive
	fmt.Println("String:", "banana" == "banana")

	fmt.Println(3 != 2)

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	fmt.Println(d1, d2)
	fmt.Println(d1 == d2)
	
	type Pessoa struct {
		Nome string
	}

	p1 := Pessoa{"Rafael"}
	p2 := Pessoa{"Rafael"}
	fmt.Println("Pessoas", p1 == p2)
}