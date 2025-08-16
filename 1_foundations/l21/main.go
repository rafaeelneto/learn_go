package main

import (
	"fmt"
	"gomod_test/matematica"

	"github.com/google/uuid"
)

func main() {
	s := matematica.Sum(10, 33)
	fmt.Printf("O valor %v \n", s)

	fmt.Println(uuid.New())
}
