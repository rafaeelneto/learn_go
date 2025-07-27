package main

const A = "Olá"

type ID int

var userId ID = 20

func main() {
	println(A)

	// local scope variable
	var a string

	println(a)

	myVarible := "my variable" // shorthand to declare varibles in this case go will infer the typing

	println(myVarible)
}
