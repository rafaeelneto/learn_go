package main

import (
	"fmt"
	"reflect"
)

func main() {
	// intergers
	fmt.Println(1, 2, 10000)
	fmt.Println("Literal é", reflect.TypeOf(42030))

	// unsigned (only positive numbers) uint8 uint16 uint32 uint64
	var b byte = 255 // alias of uint8
	fmt.Println(b)

	// signed (postive and negatives) int8 int16 int32 int64

	// int is int32 or int64 depending on the mechine x32 or x64

	// i2 int32 value that maps to unicode table of i2 character
	var i2 rune = 'n'
	fmt.Println(i2)

	// real numbers (float32, float64 -- default type literal)

	// boolean types
	bo := true;
	fmt.Println("O tipo boolean é", reflect.TypeOf(bo))

	s1 := "string type"
	fmt.Println(s1 + "!")

	fmt.Println("string size", len(s1))

	// multiple lines string
	s2 := `
	Olá
	Meu nome 
	é
	`

	fmt.Println(s2)

	char := 'a'; // rune literal
	fmt.Println("O tipo de char é", reflect.TypeOf(char))
}