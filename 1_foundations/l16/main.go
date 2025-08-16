package main

func sum(a, b int) int {
	a = 20
	return a + b
}

func sumPointer(a *int, b int) int {
	*a = 20
	return *a + b
}

func main() {
	varA := 10
	varB := 20

	println("Var A original", varA)
	println(sum(varA, varB))
	println("Var A original", varA)

	a := 10
	b := 20

	println("a original", a)
	println(sumPointer(&a, b))
	println("a original", a)
}
