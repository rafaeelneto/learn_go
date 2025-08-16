package main

func main() {
	a := 10
	println(&a)

	var ponteiro *int = &a
	println(ponteiro)

	*ponteiro = 20

	println(a)

	b := &a

	println(*b)

	c := &b
	println(c)     // address of b
	println(*c)    // adress of a
	println(*(*c)) // value of a
}
