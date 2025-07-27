package main

import "fmt"

type ID int

var userId ID = 20

func main() {

	fmt.Printf("O tipo de userId é %T", userId) // main.ID
}
