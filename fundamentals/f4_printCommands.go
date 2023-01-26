package main

import "fmt"

func main() {
	fmt.Print("mesma")
	fmt.Print("linha")

	fmt.Println(" Nova")
	fmt.Println("Linha")

	x := 3.14984

	// go doenst allow type coercion
	// fmt.Println("o valor de x é " + x)

	// ALTERNATIVAS
	xs := fmt.Sprint(x)
	fmt.Println(xs)
	fmt.Println("o valor de x é", x, "!!!")

	fmt.Printf("O valor de x formatado é %.2f", x)
	
	a := 1
	b :=2.334432
	c:= true
	d:= "opa"

	fmt.Printf("\n%d, %f, %.2f, %t, %s", a, b, b, c, d)
	fmt.Printf("\n%v %v %v %v", a, b, c, d)
}