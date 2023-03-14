package main

import "fmt"

func main() {
	numbers := [...]int{2, 4, 6, 10,505} // the three dots notations indicates that the compiler is going to count the numbers of elements on the array

	for i, number := range numbers {
		fmt.Println(i, number)
	}

	// _ ignores the index variable
	for _, number := range numbers {
		fmt.Println(	number)
	}
}