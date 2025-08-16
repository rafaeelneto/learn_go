package main

const A = "Olá"

var b bool

// global scope
var (
	c bool
	d int
	e string
)

// global scope
var (
	df bool = true
	dd int
	de string
)

func main() {
	println(A)
	println(b)

	println(c)
	println(d)
	println(e)

	// local scope variable
	var a string

	println(a)

	myVarible := "my variable" // shorthand to declare varibles in this case go will infer the typing

	println(myVarible)
}
