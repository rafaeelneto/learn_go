package main

import (
	"fmt"
	"math"
)

func main() {
	a := 2
	b := 3

	fmt.Println(a+b)
	fmt.Println(a-b)
	fmt.Println(a/b)
	fmt.Println(a*b)
	fmt.Println(a%b)

	// bitwise
	fmt.Println(a&b) // 11 & 10 = 10
	fmt.Println(a|b) // 11 | 10 = 11
	fmt.Println(a^b) // 11 ^ 10 = 01

	// operation using math
	c := 1.0
	d := 8.2
	fmt.Println(math.Max(c, d))
	fmt.Println(math.Pow(c, d))
}